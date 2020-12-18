// +build !nhooyr
// +build !raw

package wrap

import (
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var up = &websocket.Upgrader{}

func wrconn(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	wsconn, err := up.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	conn := NetConn(wsconn)
	return conn, nil
}
