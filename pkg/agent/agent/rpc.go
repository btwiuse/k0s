package agent

import (
	"bufio"
	"io"
	"log"
	"net"

	types "github.com/btwiuse/conntroll/pkg/agent"
)

var (
	_ types.RPC = (*YS)(nil)
)

func NewRPC(conn net.Conn) types.RPC {
	scanner := bufio.NewScanner(conn)
	ys := &YS{
		Scanner: scanner,
		// cmdc: make(chan Cmd),
		actions: make(chan func(types.Agent)),
		done:    make(chan struct{}),
	}
	go ys.plumbing()
	return ys
}

func (ys *YS) plumbing() {
	for ys.Scan() {
		cmd := ys.Text()
		switch cmd {
		case "PING":
			cmd = "PING"
			ys.actions <- func(ag types.Agent) {
				io.WriteString(ys.Conn, "PONG\n")
			}
		case "ACCEPT":
			cmd = "ACCEPT"
			ys.actions <- func(ag types.Agent) {
				conn, err := ag.Accept()
				if err != nil {
					log.Println(err)
				}
				// send conn to grpc server
				ag.ChanConn() <- conn
			}
		default:
			cmd = "UNKNOWN_CMD: " + cmd
		}
		log.Println(cmd)
	}
	ys.Close()
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
