package hub

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/btwiuse/conntroll/pkg/wrap"
	gows "github.com/gorilla/websocket"
	"nhooyr.io/websocket"
)

var (
	_ http.Handler = (*lys)(nil)
	_ net.Listener = (*lys)(nil)
)

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
		err    error
	)

	if gows.IsWebSocketUpgrade(r) {
		wsconn, err = websocket.Accept(w, r, &websocket.AcceptOptions{
			InsecureSkipVerify: true,
		})
		conn = websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary)
	} else {
		conn, err = wrap.Hijack(w)
	}

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
