// +build gorilla

package dialer

import (
	"net"
	"path"

	"github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/wrap"
	"github.com/gorilla/websocket"
)

var (
	_ agent.Dialer = (*gorillaDialer)(nil)
)

func New(c agent.Config) agent.Dialer {
	return &gorillaDialer{
		c: c,
	}
}

type gorillaDialer struct {
	c agent.Config
}

func (d *gorillaDialer) Dial(p string) (conn net.Conn, err error) {
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

	wsconn, _, err := websocket.DefaultDialer.Dial(u, nil)

	if err != nil {
		return nil, err
	}

	return wrap.NetConn(wsconn), nil
}
