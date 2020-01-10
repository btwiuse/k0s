package agent

import (
	"net"
	"net/http"

	"google.golang.org/grpc"
	types "k0s.io/conntroll/pkg/agent"
	"k0s.io/conntroll/pkg/agent/tty"
	"k0s.io/conntroll/pkg/api"
	"k0s.io/conntroll/pkg/api/grpcimpl"
	"k0s.io/conntroll/pkg/exporter"
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
		TtyFactory:     tty.NewFactory(cmd),
		FileServer:     http.FileServer(http.Dir("/")),
		MetricsHandler: exporter.NewHandler(),
	})
	grpcListener := NewLys()
	go grpcServer.Serve(grpcListener)
	// return grpcListener.ChanConn()
	return grpcListener
}
