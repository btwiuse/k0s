// +build gorilla
// +build !nhooyr
// +build !raw

package dialer

import (
	"net"
	"net/url"

	"github.com/gorilla/websocket"
	"k0s.io/conntroll/pkg/wrap"
)

func (d *dialr) Dial(p string) (conn net.Conn, err error) {
	var (
		c  = d.c
		ub = &url.URL{
			Scheme: c.GetSchemeWS(),
			Host:   c.GetAddr(),
			Path:   p,
		}
		u string
	)

	wsconn, _, err := websocket.DefaultDialer.Dial(u, nil)

	if err != nil {
		return nil, err
	}

	return wrap.NetConn(wsconn), nil
}
