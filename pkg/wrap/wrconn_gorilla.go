//go:build gorilla && !nhooyr

package wrap

import (
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var up = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	wsconn, err := up.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	conn := NetConn(wsconn)
	return conn, nil
}
