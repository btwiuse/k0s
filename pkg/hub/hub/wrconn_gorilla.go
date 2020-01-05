// +build gorilla

package hub

import (
	"net"
	"net/http"

	"github.com/btwiuse/conntroll/pkg/wrap"
	"github.com/gorilla/websocket"
)

var up = &websocket.Upgrader{}

func wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	wsconn, err := up.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	conn := wrap.NetConn(wsconn)
	return conn, nil
}
