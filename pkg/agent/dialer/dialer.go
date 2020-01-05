// +build !gorilla

package dialer

import (
	"context"
	"net"
	"path"

	"github.com/btwiuse/conntroll/pkg/agent"
	"nhooyr.io/websocket"
)

var (
	_ agent.Dialer = (*nhooyrDialer)(nil)
)

func New(c agent.Config) agent.Dialer {
	return &nhooyrDialer{
		c: c,
	}
}

type nhooyrDialer struct {
	c agent.Config
}

func (d *nhooyrDialer) Dial(p string) (conn net.Conn, err error) {
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
