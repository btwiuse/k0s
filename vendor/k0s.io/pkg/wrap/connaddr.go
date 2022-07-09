package wrap

import (
	"net"
)

// attach ip info to net.Conn produced by websocket.NetConn
func ConnWithAddr(conn net.Conn, addr net.Addr) net.Conn {
	return &connAddr{conn, addr}
}

func (ca *connAddr) RemoteAddr() net.Addr {
	return ca.Addr
}

type connAddr struct {
	net.Conn
	net.Addr
}

func NewAddr(network, hostport string) net.Addr {
	return &addr{network, hostport}
}

// addr implements net.Addr
type addr struct{ network, hostport string }

func (a *addr) Network() string { return a.network }

func (a *addr) String() string { return a.hostport }
