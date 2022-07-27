package server

import (
	"net"
	"net/url"

	"k0s.io/pkg/agent"
	"k0s.io/pkg/dial"
)

type dialer struct {
	c agent.Config
}

func (d *dialer) Dial(p string, q string) (conn net.Conn, err error) {
	u := &url.URL{
		Scheme:   d.c.GetSchemeWS(),
		Host:     d.c.GetAddr(),
		Path:     p,
		RawQuery: q,
	}

	return dial.Dial(u)
}
