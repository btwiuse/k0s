//go:build !gorilla && !raw
// +build !gorilla,!raw

package wrap

import (
	"net"
	"net/http"

	"k0s.io"
	"nhooyr.io/websocket"
)

func wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	wsconn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, err
	}
	wsconn.SetReadLimit(k0s.MAX_WS_MESSAGE)
	conn := NetConn(wsconn)
	addr := NewAddr("websocket", r.RemoteAddr)
	conn = ConnWithAddr(conn, addr)
	return conn, nil
}
