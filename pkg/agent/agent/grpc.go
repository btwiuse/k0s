package agent

import (
	"net"

	types "github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/agent/tty"
	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/conntroll/pkg/api/grpcimpl"
	"google.golang.org/grpc"
)

var (
	_ types.GRPCServer = (*lys)(nil)
)

// lys implements net.Listener
type lys struct {
	conns chan net.Conn
}

func (l *lys) ChanConn() chan<- net.Conn {
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

func StartGRPCServer(cmd []string) types.GRPCServer {
	grpcServer := grpc.NewServer()
	api.RegisterSessionServer(grpcServer, &grpcimpl.Session{
		TtyFactory: tty.NewFactory(cmd),
	})
	grpcListener := NewLys()
	go grpcServer.Serve(grpcListener)
	// return grpcListener.ChanConn()
	return grpcListener
}
