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

func (d *dialer) Dial(p string, h http.Header) (conn net.Conn, err error) {
	var (
		c  = d.c
		ub = &url.URL{
			Scheme: c.GetSchemeWS(),
			Host:   c.GetAddr(),
			Path:   p,
		}
		u = ub.String()
		t = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: c.GetInsecure(),
				},
			},
		}
		opts = &websocket.DialOptions{
			HTTPClient: t,
			HTTPHeader: h,
		}
	)

	wsconn, _, err := websocket.Dial(context.Background(), u, opts)
	if err != nil {
		return nil, err
	}

	return websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}
