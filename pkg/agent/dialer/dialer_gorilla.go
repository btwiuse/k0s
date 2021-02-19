// +build !nhooyr
// +build !raw

package dialer

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"k0s.io/pkg/wrap"
)

func (d *dialr) Dial(p string, q string) (conn net.Conn, err error) {
	var (
		c  = d.c
		ub = &url.URL{
			Scheme:   c.GetSchemeWS(),
			Host:     c.GetAddr(),
			Path:     p,
			RawQuery: q,
		}
		u                    = ub.String()
		wd *websocket.Dialer = websocket.DefaultDialer
	)

	if c.GetInsecure() {
		wd = &websocket.Dialer{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				RootCAs:            nil,
				InsecureSkipVerify: true,
			},
		}
	}

	wsconn, _, err := wd.Dial(u, nil)
	if err != nil {
		return nil, err
	}

	return wrap.NetConn(wsconn), nil
}
