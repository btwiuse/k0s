package tcp

import (
	"net"

	"k0s.io/conntroll/pkg/agent"
)

func NewFactory() agent.TcpFactory {
	return &factory{
		Dialer: &net.Dialer{},
	}
}

var (
	_ agent.TcpFactory = (*factory)(nil)
)

type factory struct {
	Dialer *net.Dialer
}

func (f *factory) NewProxy(addr string) (net.Conn, error) {
	return f.Dialer.Dial("tcp", addr)
}
