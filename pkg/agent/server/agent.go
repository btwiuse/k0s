package server

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"os/exec"
	"strings"

	"github.com/btwiuse/pretty"
	"golang.org/x/sync/errgroup"
	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
)

type TunnelFn = func(agent.Config) chan net.Conn

var (
	_       agent.Agent             = (*server)(nil)
	Tunnels map[api.Tunnel]TunnelFn = map[api.Tunnel]TunnelFn{}
)

type server struct {
	*errgroup.Group
	*dialer
	agent.Config
	Tunnels map[api.Tunnel]chan net.Conn
	// agent.RPC

	id   string
	name string
}

func NewAgent(c agent.Config) agent.Agent {
	var (
		eg, _   = errgroup.WithContext(context.Background())
		id      = c.GetID()
		name    = c.GetName()
		shell   = "bash"
		dialer  = &dialer{c}
		tunnels = map[api.Tunnel]chan net.Conn{}
	)

	if _, err := exec.LookPath(shell); err != nil {
		shell = "sh"
	}

	for k, v := range Tunnels {
		tunnels[k] = v(c)
	}

	return &server{
		Group:   eg,
		Config:  c,
		dialer:  dialer,
		Tunnels: tunnels,
		id:      id,
		name:    name,
	}
}

func (ag *server) TunnelChan(tun api.Tunnel) chan net.Conn {
	return ag.Tunnels[tun]
}

func (ag *server) Accept(tun api.Tunnel) (net.Conn, error) {
	var (
		conn  net.Conn
		err   error
		path  = "/api/" + strings.ToLower(tun.String())
		query = "id=" + ag.GetID()
	)

	conn, err = ag.Dial(path, query)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (ag *server) AgentRegister(conn net.Conn) (agent.RPC, error) {
	agentInfo := ag.Config.String()
	if ag.Config.GetVerbose() {
		log.Print(pretty.JSONStringLine(ag.Config))
	}
	_, err := io.WriteString(conn, agentInfo)
	if err != nil {
		return nil, err
	}

	return NewRPC(conn), nil
}

func (ag *server) Serve(rpc agent.RPC) error {
	for {
		select {
		case f := <-rpc.Actions():
			go f(ag)
		case <-rpc.Done():
			goto exit
		}
	}
exit:
	return errors.New("agent: connection to hub closed")
}

func (ag *server) ConnectAndServe() error {
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

	return errors.New("agent: serve failed")
}
