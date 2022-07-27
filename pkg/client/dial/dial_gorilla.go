//go:build gorilla && !nhooyr && !raw
// +build gorilla,!nhooyr,!raw

package dial

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"k0s.io/pkg/wrap"
)

func (d *dialer) Dial(p string, h http.Header) (conn net.Conn, err error) {
	var (
		c  = d.c
		ub = &url.URL{
			Scheme: c.GetSchemeWS(),
			Host:   c.GetAddr(),
			Path:   p,
		}
		u                    = ub.String()
		wd *websocket.Dialer = websocket.DefaultDialer
	)

	if c.GetInsecure() {
		wd = &websocket.Dialer{
			TLSClientConfig: &tls.Config{
				RootCAs:            nil,
				InsecureSkipVerify: true,
			},
		}
	}

	wsconn, _, err := wd.Dial(u, h)
	if err != nil {
		return nil, err
	}

	return wrap.NetConn(wsconn), nil
}
