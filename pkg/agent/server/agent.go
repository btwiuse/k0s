package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"

	"github.com/btwiuse/pretty"
	"golang.org/x/sync/errgroup"
	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
)

type ChannelFn = func(agent.Config) chan net.Conn

var (
	_ agent.Agent = (*server)(nil)
)

type server struct {
	*errgroup.Group
	*dialer
	agent.Config
	protocolHandlers map[api.ProtocolID]chan net.Conn
	// agent.Session

	id   string
	name string
}

func NewAgent(c agent.Config) agent.Agent {
	var (
		eg, _            = errgroup.WithContext(context.Background())
		id               = c.GetID()
		name             = c.GetName()
		shell            = "bash"
		dialer           = &dialer{c}
		protocolHandlers = map[api.ProtocolID]chan net.Conn{}
	)

	if _, err := exec.LookPath(shell); err != nil {
		shell = "sh"
	}

	ag := &server{
		Group:            eg,
		Config:           c,
		dialer:           dialer,
		protocolHandlers: protocolHandlers,
		id:               id,
		name:             name,
	}

	ag.SetProtocolHandler(api.FSID, StartFileServer)
	ag.SetProtocolHandler("fsv2", StartFileServer)
	ag.SetProtocolHandler("wsecho", StartWsEchoServer)
	ag.SetProtocolHandler("wsecho2", StartWsEcho2Server)
	ag.SetProtocolHandler("wsecho3", StartWsEcho3Server)
	ag.SetProtocolHandler(api.JsonlID, StartJsonlServer)
	ag.SetProtocolHandler(api.PingID, StartPingServer)
	ag.SetProtocolHandler(api.TerminalID, StartTerminalServer)
	ag.SetProtocolHandler(api.VersionID, StartVersionServer)
	ag.SetProtocolHandler(api.XpraID, StartXpraServer)

	return ag
}

func (ag *server) ChannelChan(p api.ProtocolID) chan net.Conn {
	_, ok := ag.protocolHandlers[p]
	if !ok {
		ag.protocolHandlers[p] = make(chan net.Conn)
	}
	return ag.protocolHandlers[p]
}

func (ag *server) AcceptProtocol(p api.ProtocolID) (net.Conn, error) {
	println("AcceptProtocol", string(p))
	var (
		conn  net.Conn
		err   error
		path  = "/api/upgrade"
		query = fmt.Sprintf("id=%s&protocol=%s", ag.GetID(), string(p))
	)

	conn, err = ag.Dial(path, query)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (ag *server) AgentRegister(conn net.Conn) (agent.Session, error) {
	agentInfo := ag.Config.String()

	if ag.Config.GetVerbose() {
		log.Print(pretty.JSONStringLine(ag.Config))
	}

	_, err := io.WriteString(conn, agentInfo)
	if err != nil {
		return nil, err
	}

	cs := NewClientSession(conn)

	return cs, nil
}

func (ag *server) Serve(cs agent.Session) error {
	for {
		select {
		case f := <-cs.Actions():
			go f(ag)
		case <-cs.Done():
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
		path = "/api/upgrade"
		// unused: "id=" + ag.GetID()
		query = ""
	)

	conn, err = ag.Dial(path, query)
	if err != nil {
		return err
	}

	cs, err := ag.AgentRegister(conn)
	if err != nil {
		return err
	}

	err = ag.Serve(cs)
	if err != nil {
		return err
	}

	return errors.New("agent: serve failed")
}

func (ag *server) SetProtocolHandler(p api.ProtocolID, fn ChannelFn) {
	ag.protocolHandlers[p] = fn(ag.Config)
}
