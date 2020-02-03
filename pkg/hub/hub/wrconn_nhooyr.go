// +build !gorilla
// +build !raw

package hub

import (
	"net"
	"net/http"

	"k0s.io/k0s/pkg/wrap"
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
