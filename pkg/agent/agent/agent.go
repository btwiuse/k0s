package agent

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/rpc"
	"os/exec"

	types "github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/agent/tty"
	"github.com/btwiuse/conntroll/pkg/api/rpcimpl"

	// "github.com/btwiuse/conntroll/pkg/uuid"
	"golang.org/x/sync/errgroup"
)

var (
	_ types.Agent = (*agent)(nil)
)

type agent struct {
	*errgroup.Group
	types.TtyFactory

	id   string
	name string
	c    types.Config
}

func NewAgent(c types.Config) types.Agent {
	eg, _ := errgroup.WithContext(context.Background())
	id := c.ID()
	name := c.Name()
	if c.Verbose() {
		log.Println("new agent", id, name)
	}
	shell := "bash"
	if _, err := exec.LookPath(shell); err != nil {
		shell = "sh"
	}
	return &agent{
		Group:      eg,
		id:         id,
		name:       name,
		TtyFactory: tty.NewFactory(c.Cmd()),
		c:          c,
	}
}

func (ag *agent) ConnectAndServe() error {
	conn, err := ag.Connect()
	if err != nil {
		return err
	}
	return ag.Serve(conn)
}

func (ag *agent) CreateSession() (net.Conn, error) {
	var (
		conn net.Conn
		err  error
	)

	conn, err = ag.Dial()
	if err != nil {
		return nil, err
	}

	_, err = conn.Write(ag.c.NewSessionRequestBody())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (ag *agent) Dial() (conn net.Conn, err error) {
	var c = ag.c
	switch c.Scheme() {
	case "http":
		conn, err = net.Dial("tcp", c.Addr())
		if err != nil {
			return nil, err
		}
		return conn, nil
	case "https":
		conn, err = tls.Dial("tcp", c.Addr(), &tls.Config{
			InsecureSkipVerify: c.Insecure(),
		})
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
	return nil, errors.New("unknown scheme")
}

func (ag *agent) Connect() (net.Conn, error) {
	var (
		conn net.Conn
		err  error
	)

	conn, err = ag.Dial()
	if err != nil {
		return nil, err
	}

	_, err = conn.Write(ag.c.NewAgentRequestBody())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (ag *agent) Serve(conn net.Conn) error {
	rpcServer := rpc.NewServer()
	rpcServer.Register(&rpcimpl.Session{Agent: ag})
	rpcServer.Register(&rpcimpl.RPC{Agent: ag})
	rpcServer.ServeConn(conn)
	return nil
}
