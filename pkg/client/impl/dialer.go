package impl

import (
	"net"
	"net/url"

	"github.com/btwiuse/wsdial"
	"k0s.io/pkg/client/config"
)

type dialer struct {
	c *config.Config
}

func (d *dialer) Dial(p string, userinfo *url.Userinfo) (conn net.Conn, err error) {
	u := &url.URL{
		Scheme: d.c.GetSchemeWS(),
		Host:   d.c.GetAddr(),
		Path:   p,
		User:   userinfo,
	}

	return wsdial.Dial(u)
}
