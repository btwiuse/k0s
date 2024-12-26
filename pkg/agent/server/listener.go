package server

import (
	"net"
)

// ChannelListener implements net.Listener
type ChannelListener struct {
	Conns chan net.Conn
}

func (cl *ChannelListener) Accept() (net.Conn, error) {
	return <-cl.Conns, nil
}

func (cl *ChannelListener) Close() error {
	// TODO: close Conns
	return nil
}

func (cl *ChannelListener) Addr() net.Addr {
	return cl
}

func (cl *ChannelListener) Network() string {
	return "hijack"
}

func (cl *ChannelListener) String() string {
	return cl.Network()
}

func NewChannelListener() *ChannelListener {
	return &ChannelListener {
		Conns: make(chan net.Conn),
	}
}
