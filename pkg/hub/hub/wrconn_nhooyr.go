// +build !gorilla

package hub

import (
	"net"
	"net/http"

	"github.com/btwiuse/conntroll/pkg/wrap"
	"nhooyr.io/websocket"
)

func wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	wsconn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, err
	}
	conn := wrap.NetConn(wsconn)
	addr := NewAddr("websocket", r.RemoteAddr)
	conn = ConnWithAddr(conn, addr)
	return conn, nil
}

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

func NewAddr(network, string string) net.Addr {
	return &addr{
		network: network,
		string:  string,
	}
}

// addr implements net.Addr
type addr struct {
	network string
	string  string
}

func (a *addr) Network() string {
	return a.network
}

func (a *addr) String() string {
	return a.string
}
