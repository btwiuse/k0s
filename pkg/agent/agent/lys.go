package agent

import (
	"net"

	types "k0s.io/conntroll/pkg/agent"
)

var (
	_ types.GrpcServer   = (*lys)(nil)
	_ types.Socks5Server = (*lys)(nil)
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
