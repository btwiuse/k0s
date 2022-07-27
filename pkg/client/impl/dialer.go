package impl

import (
	"net"
	"net/url"

	"k0s.io/pkg/client"
	"k0s.io/pkg/dial"
)

type dialer struct {
	c client.Config
}

func (d *dialer) Dial(p string, userinfo *url.Userinfo) (conn net.Conn, err error) {
	u := &url.URL{
		Scheme: d.c.GetSchemeWS(),
		Host:   d.c.GetAddr(),
		Path:   p,
		User:   userinfo,
	}

	return dial.Dial(u)
}
