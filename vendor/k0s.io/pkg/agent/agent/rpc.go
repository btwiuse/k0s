package agent

import (
	"bufio"
	"log"
	"net"
	"time"

	types "k0s.io/pkg/agent"
	"k0s.io/pkg/api"
)

var (
	_ types.RPC = (*YS)(nil)
)

func NewRPC(conn net.Conn) types.RPC {
	scanner := bufio.NewScanner(conn)
	ys := &YS{
		Conn:    conn,
		Scanner: scanner,
		actions: make(chan func(types.Agent)),
		done:    make(chan struct{}),
	}
	go ys.plumbing()
	return ys
}

func (rpc *YS) plumbing() {
	defer rpc.Close()
	for rpc.Scan() {
		cmd := rpc.Text()
		// log.Println(cmd)
		tun, err := api.FromString(cmd)
		if err != nil {
			log.Println(err)
			continue
		}
		if tun == api.Ping {
			continue
		}
		rpc.actions <- func(ag types.Agent) {
			var (
				conn net.Conn
				err  error
			)
			// make sure conn is not nil
			for i := 0; ; i++ {
				conn, err = ag.Accept(tun)
				if err != nil {
					log.Println(i, err)
					// retry on exponential interval
					time.After(time.Duration(1<<i) * time.Millisecond)
					continue
				}
				break
			}
			ag.TunnelChan(tun) <- conn
		}
	}
}

type YS struct {
	net.Conn
	*bufio.Scanner
	// cmdc chan Cmd
	actions chan func(types.Agent)
	done    chan struct{}
}

func (ys *YS) Actions() <-chan func(types.Agent) {
	return ys.actions
}

func (ys *YS) Close() {
	close(ys.done)
}

func (ys *YS) Done() <-chan struct{} {
	return ys.done
}
