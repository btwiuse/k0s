package socksdialer

import (
	"context"
	"net"

	"k0s.io/conntroll/pkg/client"
	"k0s.io/conntroll/pkg/client/wsdialer"
	"k0s.io/conntroll/pkg/socks"
)

func New(c client.Config, ep string) client.Dialer {
	wsd := wsdialer.New(c)
	d := &socksdialer{
		sd: &socks.Dialer{
			ProxyDial: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return wsd.Dial(ep)
			},
		},
	}
	return d
}

type socksdialer struct {
	sd *socks.Dialer
}

func (d *socksdialer) Dial(addr string) (net.Conn, error) {
	return d.sd.Dial("tcp", addr)
}
