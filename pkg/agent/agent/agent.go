package agent

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
	types "k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/agent/dialer"
	"k0s.io/k0s/pkg/api"
)

var (
	_ types.Agent = (*agent)(nil)
)

type agent struct {
	*errgroup.Group
	types.Config
	types.Dialer
	Tunnels map[api.Tunnel]chan net.Conn
	// types.RPC

	id   string
	name string
}

func NewAgent(c types.Config) types.Agent {
	var (
		eg, _ = errgroup.WithContext(context.Background())
		id    = c.GetID()
		name  = c.GetName()
		shell = "bash"
	)

	if _, err := exec.LookPath(shell); err != nil {
		shell = "sh"
	}

	log.Println("connected as", name)

	return &agent{
		Group:  eg,
		Config: c,
		Dialer: dialer.New(c),
		Tunnels: map[api.Tunnel]chan net.Conn{
			api.FS:       StartFileServer(c),
			api.Socks5:   StartSocks5Server(c),
			api.Redir:    StartRedirectServer(c),
			api.Metrics:  StartMetricsServer(c),
			api.Terminal: StartTerminalServer(c),
			api.Ping:     StartPingServer(c),
			api.Version:  StartVersionServer(c),
		},
		id:   id,
		name: name,
	}
}

func (ag *agent) TunnelChan(tun api.Tunnel) chan net.Conn {
	return ag.Tunnels[tun]
}

func (ag *agent) Accept(tun api.Tunnel) (net.Conn, error) {
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
