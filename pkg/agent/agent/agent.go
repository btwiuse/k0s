package agent

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"log"
	"net"
	"os/exec"

	types "github.com/btwiuse/conntroll/pkg/agent"
	"golang.org/x/net/proxy"
	"golang.org/x/sync/errgroup"
)

var (
	_ types.Agent = (*agent)(nil)
)

type agent struct {
	*errgroup.Group
	types.Config
	// types.RPC

	types.GRPCServer
	// grpcln chan<- net.Conn
	id   string
	name string
}

func NewAgent(c types.Config) types.Agent {
	eg, _ := errgroup.WithContext(context.Background())
	id := c.GetID()
	name := c.GetName()
	if c.GetVerbose() {
		log.Println("new agent", id, name)
	}
	shell := "bash"
	if _, err := exec.LookPath(shell); err != nil {
		shell = "sh"
	}
	grpcServer := StartGRPCServer(c.GetCmd())

	return &agent{
		Group:      eg,
		Config:     c,
		GRPCServer: grpcServer,
		id:         id,
		name:       name,
	}
}

func (ag *agent) Accept() (net.Conn, error) {
	var (
		conn net.Conn
		err  error
	)

	conn, err = ag.Dial()
	if err != nil {
		return nil, err
	}

	_, err = conn.Write(ag.Config.FakeHeader("/api/grpc?id=" + ag.GetID()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (ag *agent) Dial() (conn net.Conn, err error) {
	var c = ag.Config
	dialer := &net.Dialer{}
	proxyDialer := proxy.FromEnvironmentUsing(dialer)
	switch c.GetScheme() {
	case "http":
		// conn, err = net.Dial("tcp", c.GetAddr())
		conn, err = proxyDialer.Dial("tcp", c.GetAddr())
		if err != nil {
			return nil, err
		}
		return conn, nil
	case "https":
		conn, err = tls.Dial("tcp", c.GetAddr(), &tls.Config{
			// conn, err = tls.DialWithDialer(proxyDialer, "tcp", c.GetAddr(), &tls.Config{
			InsecureSkipVerify: c.GetInsecure(),
		})
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
	return nil, errors.New("unknown scheme")
}

func (ag *agent) AgentRegister(conn net.Conn) (types.RPC, error) {
	// connect
	_, err := conn.Write(ag.Config.FakeHeader("/api/yrpc"))
	if err != nil {
		return nil, err
	}

	// connect2
	_, err = io.WriteString(conn, ag.Config.String())
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
			break
		}
	}
	return errors.New("yrpc connection closed")
}

func (ag *agent) ConnectAndServe() error {
	conn, err := ag.Dial()
	if err != nil {
		return err
	}

	rpc, err := ag.AgentRegister(conn)
	if err != nil {
		return err
	}

	log.Println("Connect ok")

	err = ag.Serve(rpc)
	if err != nil {
		log.Println("Serve:", err)
		return err
	}

	return nil
}
