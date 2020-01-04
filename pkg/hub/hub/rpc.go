package hub

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	types "github.com/btwiuse/conntroll/pkg/hub"
	"github.com/btwiuse/conntroll/pkg/hub/agent"
	"github.com/btwiuse/conntroll/pkg/hub/agent/info"
	"github.com/btwiuse/conntroll/pkg/rng"
	"github.com/btwiuse/conntroll/pkg/uuid"
)

var (
	_ types.RPC = (*YS)(nil)
)

func ToRPC(conn net.Conn) types.RPC {
	rpc := &YS{
		id:            uuid.New(),
		name:          rng.New(),
		actions:       make(chan func(types.Hub)),
		created:       time.Now(),
		Conn:          conn,
		done:          make(chan struct{}),
		closeOnceDone: &sync.Once{},
		Scanner:       bufio.NewScanner(conn),
	}
	go rpc.plumbing()
	return rpc
}

func (rpc *YS) plumbing() {
	defer rpc.Close()
	for rpc.Scan() {
		cmd := rpc.Text()
		switch {
		case cmd == "PONG":
			// infinite ping/pong loop
			// rpc.Ping()
		// Register Agent
		case strings.HasPrefix(cmd, "{"):
			rpc.actions <- func(h types.Hub) {
				ifo, err := info.Decode([]byte(cmd))
				if err != nil {
					log.Println(err)
					return
				}

				var (
					id   = ifo.GetID()
					name = ifo.GetName()
				)

				if h.Has(id) {
					io.WriteString(rpc.Conn, "Error: duplicate id\n")
					return
				}

				// only set the values after we make sure !h.Has(id)
				// otherwise Unregister(h) will remove existing one
				// see func Unregister()

				rpc.agentId = id
				rpc.agentName = name

				ag := agent.NewAgent(rpc, ifo)

				h.Add(ag)
			}
		default:
			cmd = "UNKNOWN_CMD: " + cmd
			log.Println(cmd)
		}
	}
}

func (ys *YS) Actions() <-chan func(types.Hub) {
	return ys.actions
}

func (ys *YS) Unregister(h types.Hub) {
	log.Println("Disconnected:", ys.agentId, ys.agentName)
	h.Del(ys.agentId)
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
	agentId   string
	agentName string
	id        string
	name      string
	created   time.Time
	actions   chan func(types.Hub)
	net.Conn
	*bufio.Scanner
	done          chan struct{}
	closeOnceDone *sync.Once
}

func (ys *YS) NewSession() {
	_, err := io.WriteString(ys.Conn, fmt.Sprintln("ACCEPT"))
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
