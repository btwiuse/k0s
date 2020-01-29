package agent

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"os/exec"

	"github.com/btwiuse/pretty"
	"golang.org/x/sync/errgroup"
	types "k0s.io/conntroll/pkg/agent"
	"k0s.io/conntroll/pkg/agent/dialer"
)

var (
	_ types.Agent = (*agent)(nil)
)

type agent struct {
	*errgroup.Group
	types.Config
	types.Dialer
	// types.RPC

	types.FileServer
	types.GrpcServer // deprecated in favor of types.TerminalServer
	types.Socks5Server
	types.MetricsServer
	types.TerminalServer
	// grpcln chan<- net.Conn
	id   string
	name string
}

func NewAgent(c types.Config) types.Agent {
	var (
		eg, _ = errgroup.WithContext(context.Background())
		id    = c.GetID()
		name  = c.GetName()
		shell = "bash"

		fileServer     = StartFileServer(c)
		grpcServer     = StartGrpcServer(c)
		socks5Server   = StartSocks5Server(c)
		metricsServer  = StartMetricsServer(c)
		terminalServer = StartTerminalServer(c)
	)
	if c.GetVerbose() {
		log.Println("new agent", id, name)
	}
	if _, err := exec.LookPath(shell); err != nil {
		shell = "sh"
	}

	return &agent{
		Group:          eg,
		Config:         c,
		Dialer:         dialer.New(c),
		FileServer:     fileServer,
		GrpcServer:     grpcServer,
		Socks5Server:   socks5Server,
		MetricsServer:  metricsServer,
		TerminalServer: terminalServer,
		id:             id,
		name:           name,
	}
}

func (ag *agent) FSChanConn() chan<- net.Conn {
	return ag.FileServer.ChanConn()
}

func (ag *agent) GrpcChanConn() chan<- net.Conn {
	return ag.GrpcServer.ChanConn()
}

func (ag *agent) Socks5ChanConn() chan<- net.Conn {
	return ag.Socks5Server.ChanConn()
}

func (ag *agent) MetricsChanConn() chan<- net.Conn {
	return ag.MetricsServer.ChanConn()
}

func (ag *agent) TerminalChanConn() chan<- net.Conn {
	return ag.TerminalServer.ChanConn()
}

func (ag *agent) AcceptFS() (net.Conn, error) {
	var (
		conn  net.Conn
		err   error
		path  = "/api/fs"
		query = "id=" + ag.GetID()
	)

	conn, err = ag.Dial(path, query)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (ag *agent) AcceptSocks5() (net.Conn, error) {
	var (
		conn  net.Conn
		err   error
		path  = "/api/socks5"
		query = "id=" + ag.GetID()
	)

	conn, err = ag.Dial(path, query)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (ag *agent) AcceptGrpc() (net.Conn, error) {
	var (
		conn  net.Conn
		err   error
		path  = "/api/grpc"
		query = "id=" + ag.GetID()
	)

	conn, err = ag.Dial(path, query)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (ag *agent) AcceptMetrics() (net.Conn, error) {
	var (
		conn  net.Conn
		err   error
		path  = "/api/metrics"
		query = "id=" + ag.GetID()
	)

	conn, err = ag.Dial(path, query)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (ag *agent) AcceptTerminal() (net.Conn, error) {
	var (
		conn  net.Conn
		err   error
		path  = "/api/terminal"
		query = "id=" + ag.GetID()
	)

	conn, err = ag.Dial(path, query)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (ag *agent) AgentRegister(conn net.Conn) (types.RPC, error) {
	agentInfo := ag.Config.String()
	if ag.Config.GetVerbose() {
		log.Print(pretty.JSONString(ag.Config))
	}
	_, err := io.WriteString(conn, agentInfo)
	if err != nil {
		return nil, err
	}

	return NewRPC(conn), nil
}

func (ag *agent) Serve(rpc types.RPC) error {
	for {
		select {
		case f := <-rpc.Actions():
			go f(ag)
		case <-rpc.Done():
			goto exit
		}
	}
exit:
	return errors.New("rpc connection closed")
}

func (ag *agent) ConnectAndServe() error {
	var (
		conn net.Conn
		err  error
		path = "/api/rpc"
		// unused: "id=" + ag.GetID()
		query = ""
	)

	conn, err = ag.Dial(path, query)
	if err != nil {
		return err
	}

	rpc, err := ag.AgentRegister(conn)
	if err != nil {
		return err
	}

	err = ag.Serve(rpc)
	if err != nil {
		return err
	}

	return nil
}
