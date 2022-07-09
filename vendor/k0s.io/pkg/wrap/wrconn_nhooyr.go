//go:build nhooyr && !gorilla && !raw
// +build nhooyr,!gorilla,!raw

package wrap

import (
	"net"
	"net/http"

	"nhooyr.io/websocket"
)

func wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	wsconn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, err
	}
	conn := NetConn(wsconn)
	addr := NewAddr("websocket", r.RemoteAddr)
	conn = ConnWithAddr(conn, addr)
	return conn, nil
}
