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
	_ agent.RPC = (*YS)(nil)
)

func NewRPC(conn net.Conn) agent.RPC {
	scanner := bufio.NewScanner(conn)
	ys := &YS{
		Conn:    conn,
		Scanner: scanner,
		actions: make(chan func(agent.Agent)),
		done:    make(chan struct{}),
	}
	// go ys.plumbing()
	go ys.plumbingChan()
	return ys
}

func (rpc *YS) plumbingChan() {
	defer rpc.Close()
	for rpc.Scan() {
		cmd := rpc.Text()
		// log.Println(cmd)
		p := api.ProtocolID(cmd)
		if p == api.PingID {
			continue
		}
		rpc.actions <- func(ag agent.Agent) {
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

type YS struct {
	net.Conn
	*bufio.Scanner
	// cmdc chan Cmd
	actions chan func(agent.Agent)
	done    chan struct{}
}

func (ys *YS) Actions() <-chan func(agent.Agent) {
	return ys.actions
}

func (ys *YS) Close() {
	close(ys.done)
}

func (ys *YS) Done() <-chan struct{} {
	return ys.done
}
