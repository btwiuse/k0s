package agent

import (
	"bufio"
	"io"
	"log"
	"net"
	"time"

	types "k0s.io/conntroll/pkg/agent"
	"k0s.io/conntroll/pkg/spinner"
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
	// ys.spinner.Start()
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
				var (
					conn net.Conn
					err  error
				)
				// make sure conn is not nil
				for i := 0; ; i++ {
					conn, err = ag.AcceptGrpc()
					if err != nil {
						log.Println(i, err)
						// retry on exponential interval
						time.After(time.Duration(1<<i) * time.Millisecond)
						continue
					}
					break
				}
				// send conn to grpc server
				ag.GrpcChanConn() <- conn
			}
			// log.Println(cmd)
		case "SOCKS5":
			cmd = "SOCKS5"
			rpc.actions <- func(ag types.Agent) {
				var (
					conn net.Conn
					err  error
				)
				for i := 0; ; i++ {
					conn, err = ag.AcceptSocks5()
					if err != nil {
						log.Println(i, err)
						time.After(time.Duration(1<<i) * time.Millisecond)
						continue
					}
					break
				}
				// send conn to socks5 server
				ag.Socks5ChanConn() <- conn
			}
		case "FS":
			cmd = "FS"
			rpc.actions <- func(ag types.Agent) {
				var (
					conn net.Conn
					err  error
				)
				for i := 0; ; i++ {
					conn, err = ag.AcceptFS()
					if err != nil {
						log.Println(i, err)
						time.After(time.Duration(1<<i) * time.Millisecond)
						continue
					}
					break
				}
				// send conn to socks5 server
				ag.FSChanConn() <- conn
			}
		case "METRICS":
			cmd = "METRICS"
			rpc.actions <- func(ag types.Agent) {
				var (
					conn net.Conn
					err  error
				)
				for i := 0; ; i++ {
					conn, err = ag.AcceptMetrics()
					if err != nil {
						log.Println(i, err)
						time.After(time.Duration(1<<i) * time.Millisecond)
						continue
					}
					break
				}
				// send conn to socks5 server
				ag.MetricsChanConn() <- conn
			}
		default:
			cmd = "UNKNOWN_CMD: " + cmd
			log.Println(cmd)
			// TODO: investigate 'Error: duplicate id'
			break
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
