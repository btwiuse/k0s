package hub

import (
	"context"
	"log"
	"net"
	"net/http"

	"nhooyr.io/websocket"
)

var (
	_ http.Handler = (*lys)(nil)
	_ net.Listener = (*lys)(nil)
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

// lys is an http handler to net.Listener converter
// lys implements net.Listener/http.Handler
type lys struct {
	conns chan net.Conn
}

// support two methods of converting http request to net.Conn
// 1. websocket  (http2 support coming)
// 2. hijack     (no http2 support)
func (l *lys) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		wsconn *websocket.Conn
		conn   net.Conn
		addr   net.Addr
		err    error
	)

	wsconn, err = websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	conn = websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary)
	addr = NewAddr("websocket", r.RemoteAddr)
	conn = ConnWithAddr(conn, addr)

	if err != nil {
		log.Println("error ws accept:", err)
		return
	}

	l.conns <- conn
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

func NewHandleHijackListener() *lys {
	return &lys{conns: make(chan net.Conn)}
}
