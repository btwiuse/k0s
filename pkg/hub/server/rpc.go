package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/btwiuse/rng"
	"k0s.io/pkg/api"
	"k0s.io/pkg/hub"
	"k0s.io/pkg/hub/agent"
	"k0s.io/pkg/hub/agent/info"
)

var (
	_ hub.Session = (*ServerSession)(nil)
)

func NewServerSession(conn net.Conn) hub.Session {

	ss := &ServerSession{
		id:            "00000000-0000-0000-0000-000000000000",
		name:          rng.NewDocker(),
		actions:       make(chan func(hub.Hub), 1),
		created:       time.Now(),
		Conn:          conn,
		done:          make(chan struct{}),
		closeOnceDone: &sync.Once{},
		Scanner:       bufio.NewScanner(conn),
	}

	ss.populate()

	go ss.plumbing()

	return ss
}

func (ss *ServerSession) populate() {
	ss.Scan()
	cmd := ss.Text()

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

	ss.id = id
	ss.name = name

	ss.actions <- func(h hub.Hub) {
		// clobber previous connection, if any
		if h.Has(id) {
			h.Del(id)
		}

		ag := agent.NewAgent(ss, ifo)

		h.Add(ag)
	}
}

func (ss *ServerSession) plumbing() {
	defer func() {
		ss.Close()
		// println("hub close")
	}()
	for ss.Scan() {
		cmd := ss.Text()
		switch {
		case cmd == "PONG":
			// infinite ping/pong loop
			// ss.Ping()
		default:
			// cmd = "UNKNOWN_CMD: " + cmd
			// log.Println(cmd)
		}
	}
}

func (ss *ServerSession) Actions() <-chan func(hub.Hub) {
	return ss.actions
}

func (ss *ServerSession) Close() {
	ss.closeOnceDone.Do(func() {
		close(ss.done)
	})
}

func (ss *ServerSession) Done() <-chan struct{} {
	return ss.done
}

func (ss *ServerSession) Time() time.Time {
	return ss.created
}

func (ss *ServerSession) Name() string {
	return ss.name
}

func (ss *ServerSession) ID() string {
	return ss.id
}

func (ss *ServerSession) RemoteIP() string {
	remote_hostport := ss.Conn.RemoteAddr().String()
	if !strings.Contains(remote_hostport, ":") {
		return remote_hostport
	}
	ip, _, _ := net.SplitHostPort(remote_hostport)
	return ip
}

type ServerSession struct {
	id      string
	name    string
	created time.Time
	actions chan func(hub.Hub)
	net.Conn
	*bufio.Scanner
	done          chan struct{}
	closeOnceDone *sync.Once
}

func (ss *ServerSession) OpenChannel(p api.ProtocolID) {
	cmd := string(p)
	_, err := io.WriteString(ss.Conn, fmt.Sprintln(cmd))
	if err != nil {
		ss.Close()
	}
}

func (ss *ServerSession) Ping() {
	_, err := io.WriteString(ss.Conn, fmt.Sprintln("PING"))
	if err != nil {
		ss.Close()
	}
}
