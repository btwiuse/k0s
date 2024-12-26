package server

import (
	"net"
)

// lys implements net.Listener
type lys struct {
	Conns chan net.Conn
}

func (l *lys) Accept() (net.Conn, error) {
	return <-l.Conns, nil
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

func NewChannelListener() *lys {
	return &lys{
		Conns: make(chan net.Conn),
	}
}
