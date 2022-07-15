//go:build !gorilla && !raw
// +build !gorilla,!raw

package wrap

import (
	"net"
	"net/http"

	"nhooyr.io/websocket"
	"k0s.io/pkg"
)

func wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	wsconn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, err
	}
	wsconn.SetReadLimit(pkg.MAX_WS_MESSAGE)
	conn := NetConn(wsconn)
	addr := NewAddr("websocket", r.RemoteAddr)
	conn = ConnWithAddr(conn, addr)
	return conn, nil
}
