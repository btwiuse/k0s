// +build gorilla

package hub

import (
	"net"
	"net/http"

	"github.com/gorilla/websocket"
	"k0s.io/conntroll/pkg/wrap"
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
