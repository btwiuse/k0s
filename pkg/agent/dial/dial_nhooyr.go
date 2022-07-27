//go:build !gorilla && !raw
// +build !gorilla,!raw

package dial

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"

	"nhooyr.io/websocket"
)

func (d *dialer) Dial(p string, q string) (conn net.Conn, err error) {
	var (
		c  = d.c
		ub = &url.URL{
			Scheme:   c.GetSchemeWS(),
			Host:     c.GetAddr(),
			Path:     p,
			RawQuery: q,
		}
		u = ub.String()
		t = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: c.GetInsecure(),
				},
			},
		}
		opts = &websocket.DialOptions{
			HTTPClient: t,
		}
	)

	wsconn, _, err := websocket.Dial(context.Background(), u, opts)
	if err != nil {
		return nil, err
	}

	return websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}
