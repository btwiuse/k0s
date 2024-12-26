package server

import (
	"bufio"
	"log"
	"net"
	"time"

	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
)

var (
	_ agent.Session = (*ClientSession)(nil)
)

func NewClientSession(conn net.Conn) agent.Session {
	scanner := bufio.NewScanner(conn)
	cs := &ClientSession{
		Conn:    conn,
		Scanner: scanner,
		actions: make(chan func(agent.Agent)),
		done:    make(chan struct{}),
	}
	go cs.plumbing()
	return cs
}

func (cs *ClientSession) plumbing() {
	defer cs.Close()
	for cs.Scan() {
		cmd := cs.Text()
		// log.Println(cmd)
		p := api.ProtocolID(cmd)
		if p == "PING" {
			continue
		}
		cs.actions <- func(ag agent.Agent) {
			var (
				conn net.Conn
				err  error
			)
			// make sure conn is not nil
			for i := 0; ; i++ {
				conn, err = ag.AcceptProtocol(p)
				if err != nil {
					log.Println(i, err)
					// retry on exponential interval
					time.After(time.Duration(1<<i) * time.Millisecond)
					continue
				}
				break
			}
			ag.ChannelChan(p) <- conn
		}
	}
}

type ClientSession struct {
	net.Conn
	*bufio.Scanner
	// cmdc chan Cmd
	actions chan func(agent.Agent)
	done    chan struct{}
}

func (cs *ClientSession) Actions() <-chan func(agent.Agent) {
	return cs.actions
}

func (cs *ClientSession) Close() {
	close(cs.done)
}

func (cs *ClientSession) Done() <-chan struct{} {
	return cs.done
}
