//go:build js

package server

import (
	"context"
	"net"
	"net/url"

	"github.com/coder/websocket"
	"k0s.io/pkg/agent/config"
)

var wsDialOpts = &websocket.DialOptions{}

type dialer struct {
	c *config.Config
}

func (d *dialer) Dial(p string, q string) (conn net.Conn, err error) {
	u := &url.URL{
		Scheme:   d.c.GetSchemeWS(),
		Host:     d.c.GetAddr(),
		Path:     p,
		RawQuery: q,
	}

	wsconn, _, err := websocket.Dial(context.Background(), u.String(), wsDialOpts)
	if err != nil {
		return nil, err
	}

	return websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}
