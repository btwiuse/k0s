package hub

import (
	"bufio"
	"fmt"
	"io"

	// "log"
	"net"
	"sync"
	"time"

	"k0s.io/k0s/pkg/api"
	types "k0s.io/k0s/pkg/hub"
	"k0s.io/k0s/pkg/hub/agent"
	"k0s.io/k0s/pkg/hub/agent/info"
	"k0s.io/k0s/pkg/rng"
)

var (
	_ types.RPC = (*YS)(nil)
)

func ToRPC(conn net.Conn) types.RPC {

	rpc := &YS{
		id:            "00000000-0000-0000-0000-000000000000",
		name:          rng.New(),
		actions:       make(chan func(types.Hub), 1),
		created:       time.Now(),
		Conn:          conn,
		done:          make(chan struct{}),
		closeOnceDone: &sync.Once{},
		Scanner:       bufio.NewScanner(conn),
	}

	rpc.register()

	go rpc.plumbing()

	return rpc
}

func (rpc *YS) register() {
	rpc.Scan()
	cmd := rpc.Text()
	// log.Println(cmd)

	ifo, err := info.Decode([]byte(cmd))
	if err != nil {
		// TODO: notify client
		// log.Println(err)
		return
	}

	var (
		id   = ifo.GetID()
		name = ifo.GetName()
	)

	rpc.id = id
	rpc.name = name

	rpc.actions <- func(h types.Hub) {
		// clobber previous connection, if any
		if h.Has(id) {
			h.Del(id)
		}

		ag := agent.NewAgent(rpc, ifo)

		h.Add(ag)
	}
}

func (rpc *YS) plumbing() {
	defer func() {
		rpc.Close()
		// println("hub close")
	}()
	for rpc.Scan() {
		cmd := rpc.Text()
		// log.Println(cmd)
		switch {
		case cmd == "PONG":
			// infinite ping/pong loop
			// rpc.Ping()
		default:
			// cmd = "UNKNOWN_CMD: " + cmd
			// log.Println(cmd)
		}
	}
}

func (ys *YS) Actions() <-chan func(types.Hub) {
	return ys.actions
}

func (ys *YS) Close() {
	ys.closeOnceDone.Do(func() {
		close(ys.done)
	})
}

func (ys *YS) Done() <-chan struct{} {
	return ys.done
}

func (ys *YS) Time() time.Time {
	return ys.created
}

func (ys *YS) Name() string {
	return ys.name
}

func (ys *YS) ID() string {
	return ys.id
}

func (ys *YS) RemoteIP() string {
	ip, _, _ := net.SplitHostPort(ys.Conn.RemoteAddr().String())
	return ip
}

type YS struct {
	id      string
	name    string
	created time.Time
	actions chan func(types.Hub)
	net.Conn
	*bufio.Scanner
	done          chan struct{}
	closeOnceDone *sync.Once
}

func (ys *YS) NewTunnel(tun api.Tunnel) {
	cmd := tun.String()
	_, err := io.WriteString(ys.Conn, fmt.Sprintln(cmd))
	if err != nil {
		ys.Close()
	}
}

func (ys *YS) Ping() {
	_, err := io.WriteString(ys.Conn, fmt.Sprintln("PING"))
	if err != nil {
		ys.Close()
	}
}
