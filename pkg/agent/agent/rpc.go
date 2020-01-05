package agent

import (
	"bufio"
	"io"
	"log"
	"net"
	"time"

	types "github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/spinner"
)

var (
	_ types.RPC = (*YS)(nil)
)

func NewRPC(conn net.Conn) types.RPC {
	scanner := bufio.NewScanner(conn)
	pingChan := make(chan time.Time)
	ys := &YS{
		Conn:     conn,
		Scanner:  scanner,
		pingChan: pingChan,
		actions:  make(chan func(types.Agent)),
		done:     make(chan struct{}),
		spinner:  spinner.New(spinner.CharSets[9], pingChan),
	}
	ys.spinner.Start()
	go ys.plumbing()
	return ys
}

func (rpc *YS) plumbing() {
	defer rpc.Close()
	for rpc.Scan() {
		cmd := rpc.Text()
		switch cmd {
		case "PING":
			cmd = "PING"
			rpc.actions <- func(ag types.Agent) {
				err := rpc.Pong()
				if err != nil {
					return
				}
				rpc.pingChan <- time.Now()
			}
		case "ACCEPT":
			cmd = "ACCEPT"
			rpc.actions <- func(ag types.Agent) {
				conn, err := ag.Accept()
				if err != nil {
					log.Println(err)
				}
				// send conn to grpc server
				ag.ChanConn() <- conn
			}
			log.Println(cmd)
		default:
			cmd = "UNKNOWN_CMD: " + cmd
			log.Println(cmd)
		}
	}
}

func (ys *YS) Pong() error {
	_, err := io.WriteString(ys.Conn, "PONG\n")
	return err
}

type YS struct {
	net.Conn
	*bufio.Scanner
	// cmdc chan Cmd
	actions  chan func(types.Agent)
	done     chan struct{}
	spinner  *spinner.Spinner
	pingChan chan time.Time
}

func (ys *YS) Actions() <-chan func(types.Agent) {
	return ys.actions
}

func (ys *YS) Close() {
	ys.spinner.Stop()
	close(ys.done)
}

func (ys *YS) Done() <-chan struct{} {
	return ys.done
}
