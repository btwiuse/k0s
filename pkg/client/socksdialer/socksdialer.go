package socksdialer

import (
	"context"
	"net"
	"net/url"

	"k0s.io/k0s/pkg/client"
	"k0s.io/k0s/pkg/client/wsdialer"
	"k0s.io/k0s/pkg/socks"
)

func New(c client.Config, ep string) client.Dialer {
	wsd := wsdialer.New(c)
	sd := socks.NewDialer("tcp", ep)
	sd.ProxyDial = func(ctx context.Context, network, addr string) (net.Conn, error) { return wsd.Dial(ep, nil) }
	d := &socksdialer{sd: sd}
	return d
}

type socksdialer struct {
	sd *socks.Dialer
}

func (d *socksdialer) Dial(addr string, userinfo *url.Userinfo) (net.Conn, error) {
	return d.sd.Dial("tcp", addr)
}
