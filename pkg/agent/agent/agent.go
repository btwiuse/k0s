package agent

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net"

	// "net/rpc"
	"os/exec"

	types "github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/agent/tty"
	"github.com/btwiuse/conntroll/pkg/api/grpcimpl"

	// "github.com/btwiuse/conntroll/pkg/api/rpcimpl"
	// "github.com/btwiuse/conntroll/pkg/api/yrpcimpl"
	"github.com/btwiuse/conntroll/pkg/api"
	"google.golang.org/grpc"

	"github.com/yrpc/yrpc"
	// "github.com/btwiuse/conntroll/pkg/uuid"
	"golang.org/x/sync/errgroup"
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
	switch c.GetScheme() {
	case "http":
		conn, err = net.Dial("tcp", c.GetAddr())
		if err != nil {
			return nil, err
		}
		return conn, nil
	case "https":
		conn, err = tls.Dial("tcp", c.GetAddr(), &tls.Config{
			InsecureSkipVerify: c.GetInsecure(),
		})
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
	return nil, errors.New("unknown scheme")
}

func (ag *agent) YRPCConnect(conn net.Conn) (*YS, error) {
	// connect
	_, err := conn.Write(ag.Config.FakeHeader("/api/yrpc"))
	if err != nil {
		return nil, err
	}

	// confirmation
	_, err = ioutil.ReadAll(io.LimitReader(conn, 1))
	if err != nil {
		return nil, err
	}

	// dial2
	conf := yrpc.ClientConfig{
		OverlayNetwork: func(addr string, dc yrpc.DialConfig) (net.Conn, error) { return conn, nil },
	}

	yconn, err := yrpc.NewConnection("-", conf, nil)
	if err != nil {
		return nil, err
	}
	log.Println("sending api.AgentRegisterRequest")

	// connect2

	confStr := ag.Config.String()
	// log.Println(confStr)

	w, resp, err := yconn.StreamRequest(api.AgentRegisterRequest, 0, []byte(confStr))
	if err != nil {
		return nil, err
	}
	// w.StartWrite(api.AgentRegisterRequest)
	// # TODO change this part to sending client info
	// w.WriteBytes([]byte(confStr))
	// w.EndWrite(false)

	// connect2 confirm
	frame, err := resp.GetFrame()
	if err != nil {
		return nil, err
	}
	confirmMessage := string(frame.Payload)
	log.Println("confirm", confirmMessage)
	// time.Sleep(time.Hour)

	return &YS{
		StreamWriter: w,
		FrameCh:      frame.FrameCh(),
	}, nil
}

type YS struct {
	StreamWriter yrpc.StreamWriter
	FrameCh      <-chan *yrpc.Frame
	next         *yrpc.Frame
}

func (ys *YS) Next() bool {
	f := <-ys.FrameCh
	if f == nil {
		return false
	}
	ys.next = f
	return true
}

func (ys *YS) Frame() *yrpc.Frame {
	return ys.next
}

func (ys *YS) WriteCmdPayload(cmd yrpc.Cmd, payload []byte) {
	w := ys.StreamWriter
	w.StartWrite(cmd)
	w.WriteBytes(payload)
	w.EndWrite(false)
}

func (ag *agent) YRPCServe(ys *YS) error {
	for ys.Next() {
		f := ys.Frame()
		recv := string(f.Payload)
		cmd := ""
		switch f.Cmd {
		case api.Ping:
			cmd = "Ping"
			ys.WriteCmdPayload(api.Pong, []byte(recv))
		case api.AcceptRequest:
			go func() {
				conn, err := ag.Accept()
				if err != nil {
					log.Println(err)
				}
				ag.ch <- conn // GS handle connection
			}()
			cmd = "AcceptRequest"
			ys.WriteCmdPayload(api.AcceptRequest, []byte("OK"))
		default:
			cmd = "unknown cmd"
		}
		log.Println(cmd, recv)
	}
	return errors.New("yrpc connection closed")
}

func (ag *agent) YRPCConnectAndServe() error {
	conn, err := ag.Dial()
	if err != nil {
		return err
	}

	ys, err := ag.YRPCConnect(conn)
	if err != nil {
		return err
	}

	log.Println("YRPCConnect ok")

	err = ag.YRPCServe(ys)
	if err != nil {
		log.Println("YRPCServe:", err)
		return err
	}

	return nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////
/* deprecated: use YRPCConnect
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
*/

/* deprecated: use YRPCServe
func (ag *agent) Serve(conn net.Conn) error {
	rpcServer := rpc.NewServer()
	rpcServer.Register(&rpcimpl.Session{Agent: ag})
	rpcServer.Register(&rpcimpl.RPC{Agent: ag})
	rpcServer.ServeConn(conn)
	return nil
}
*/

/* deprecated: use YRPCConnectAndServe
func (ag *agent) ConnectAndServe() error {
	conn, err := ag.Connect()
	if err != nil {
		return err
	}
	return ag.Serve(conn)
}
*/

/*
func (ag *agent) Accept() (net.Conn, error) {
	<-ag.grpcConns, nil
}
*/
