package agent

import (
	"bufio"
	"context"
	"crypto/tls"
	"errors"
	"io"
	"log"
	"net"
	"os/exec"

	types "github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/agent/tty"
	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/conntroll/pkg/api/grpcimpl"
	"golang.org/x/net/proxy"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	_ types.Agent = (*agent)(nil)
)

// lys implements net.Listener
type lys struct {
	conns chan net.Conn
}

func (l *lys) Chan() chan<- net.Conn {
	return l.conns
}

func (l *lys) Accept() (net.Conn, error) {
	return <-l.conns, nil
}

func (l *lys) Close() error {
	return nil
}

func (l *lys) Addr() net.Addr {
	return l
}

func (l *lys) Network() string {
	return "hijack"
}

func (l *lys) String() string {
	return l.Network()
}

func NewLys() *lys {
	return &lys{
		conns: make(chan net.Conn),
	}
}

type agent struct {
	*errgroup.Group
	types.RPC

	types.Config
	ch   chan<- net.Conn
	id   string
	name string
}

func StartGRPCServer(cmd []string) chan<- net.Conn {
	grpcServer := grpc.NewServer()
	api.RegisterSessionServer(grpcServer, &grpcimpl.Session{
		TtyFactory: tty.NewFactory(cmd),
	})
	grpcListener := NewLys()
	go grpcServer.Serve(grpcListener)
	return grpcListener.Chan()
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
	ch := StartGRPCServer(c.GetCmd())

	return &agent{
		Group:  eg,
		id:     id,
		name:   name,
		ch:     ch,
		Config: c,
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

func (ag *agent) Connect(conn net.Conn) (*YS, error) {
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

	return &YS{
		Scanner: bufio.NewScanner(conn),
	}, nil
}

type YS struct {
	// net.Conn
	*bufio.Scanner
}

func (ag *agent) Serve(ys *YS) error {
	for ys.Scan() {
		cmd := ys.Text()
		switch cmd {
		case "ACCEPT":
			go func() {
				conn, err := ag.Accept()
				if err != nil {
					log.Println(err)
				}
				// send conn to grpc server
				ag.ch <- conn
			}()
		default:
			cmd = "UNKNOWN_CMD: " + cmd
		}
		log.Println(cmd)
	}
	return errors.New("yrpc connection closed")
}

func (ag *agent) ConnectAndServe() error {
	conn, err := ag.Dial()
	if err != nil {
		return err
	}

	ys, err := ag.Connect(conn)
	if err != nil {
		return err
	}

	log.Println("Connect ok")

	err = ag.Serve(ys)
	if err != nil {
		log.Println("Serve:", err)
		return err
	}

	return nil
}
