package agent

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/rpc"

	types "github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/agent/config"
	"github.com/btwiuse/conntroll/pkg/agent/tty"
	rpcimpl "github.com/btwiuse/conntroll/pkg/api/rpc/impl"

	// "github.com/btwiuse/conntroll/pkg/uuid"
	"golang.org/x/sync/errgroup"
)

var (
	_ types.Agent = (*agent)(nil)
)

type agent struct {
	*errgroup.Group
	types.TtyFactory

	id string
	c  *config.Config
}

func NewAgent(c *config.Config) types.Agent {
	eg, _ := errgroup.WithContext(context.Background())
	id := c.ID
	log.Println("new agent", id)
	return &agent{
		Group:      eg,
		id:         id,
		TtyFactory: tty.NewFactory([]string{"/usr/bin/env", "TERM=xterm", "bash"}),
		c:          c,
	}
}

func (ag *agent) ConnectAndServe() error {
	conn, err := ag.Connect()
	if err != nil {
		return err
	}

	rpcServer := rpc.NewServer()
	rpcServer.Register(&rpcimpl.Session{Agent: ag})
	rpcServer.Register(&rpcimpl.RPC{Agent: ag})
	rpcServer.ServeConn(conn)
	return nil
}

func (ag *agent) CreateSession() (net.Conn, error) {
	var c = ag.c
	var (
		conn net.Conn
		err  error
		port string = c.Port()
	)

	if c.Scheme == "http" {
		if port == "" {
			port = "80"
		}
		conn, err = net.Dial("tcp", c.Hostname()+":"+port)
		if err != nil {
			return nil, err
		}
	}

	if c.Scheme == "https" {
		if port == "" {
			port = "443"
		}
		conn, err = tls.Dial("tcp", c.Hostname()+":"+port, &tls.Config{
			InsecureSkipVerify: true,
		})
		if err != nil {
			return nil, err
		}
	}

	_, err = conn.Write(c.NewSessionRequestBody())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (ag *agent) Connect() (net.Conn, error) {
	var c = ag.c
	var (
		conn net.Conn
		err  error
		port string = c.Port()
	)

	if c.Scheme == "http" {
		if c.Port() == "" {
			port = "80"
		}
		conn, err = net.Dial("tcp", c.Hostname()+":"+port)
		if err != nil {
			return nil, err
		}
	}

	if c.Scheme == "https" {
		if c.Port() == "" {
			port = "443"
		}
		fmt.Println(c.Hostname() + ":" + port)
		conn, err = tls.Dial("tcp", c.Hostname()+":"+port, &tls.Config{
			InsecureSkipVerify: true,
		})
		if err != nil {
			return nil, err
		}
	}

	_, err = conn.Write(c.NewAgentRequestBody())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
