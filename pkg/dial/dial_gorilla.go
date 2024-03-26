//go:build gorilla && !nhooyr

package dial

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/btwiuse/wsconn"
)

func Dial(u *url.URL) (conn net.Conn, err error) {
	var wd *websocket.Dialer = &websocket.Dialer{
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{
			RootCAs:            nil,
			InsecureSkipVerify: true,
		},
	}

	wsconn, _, err := wd.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return wsconn.NetConn(wsconn), nil
}
