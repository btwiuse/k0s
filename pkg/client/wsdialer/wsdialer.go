package wsdialer

import (
	"context"
	"net"
	"path"

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
		c = d.c
		u string
	)

	switch c.GetScheme() {
	case "http":
		u = "ws://" + path.Join(c.GetAddr(), p)
	case "https":
		u = "wss://" + path.Join(c.GetAddr(), p)
	}

	wsconn, _, err := websocket.Dial(context.Background(), u, nil)

	if err != nil {
		return nil, err
	}

	return websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}
