//go:build js

package impl

import (
	"context"
	"net"
	"net/url"

	"k0s.io/pkg/client/config"
	"nhooyr.io/websocket"
)

type dialer struct {
	c *config.Config
}

func (d *dialer) Dial(p string, userinfo *url.Userinfo) (conn net.Conn, err error) {
	u := &url.URL{
		Scheme: d.c.GetSchemeWS(),
		Host:   d.c.GetAddr(),
		Path:   p,
	}

	wsconn, _, err := websocket.Dial(context.Background(), u.String(), nil)
	if err != nil {
		return nil, err
	}

	return websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}
