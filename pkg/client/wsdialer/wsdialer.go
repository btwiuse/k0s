package wsdialer

import (
	"context"
	"net"
	"net/url"

	"k0s.io/conntroll/pkg/client"
	"nhooyr.io/websocket"
)

var (
	_ client.Dialer = (*wsdialer)(nil)
)

func New(c client.Config) client.Dialer {
	return &wsdialer{
		c: c,
	}
}

type wsdialer struct {
	c client.Config
}

func (d *wsdialer) Dial(p string) (conn net.Conn, err error) {
	var (
		c  = d.c
		ub = &url.URL{
			Scheme: c.GetSchemeWS(),
			Host:   c.GetAddr(),
			Path:   p,
		}
		u = ub.String()
	)

	wsconn, _, err := websocket.Dial(context.Background(), u, nil)

	if err != nil {
		return nil, err
	}

	return websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}
