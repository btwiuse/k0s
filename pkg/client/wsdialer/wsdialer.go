package wsdialer

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"net"
	"net/http"
	"net/url"

	"k0s.io/k0s/pkg/client"
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

func (d *wsdialer) Dial(p string, userinfo *url.Userinfo) (conn net.Conn, err error) {
	var (
		c  = d.c
		ub = &url.URL{
			Scheme: c.GetSchemeWS(),
			Host:   c.GetAddr(),
			Path:   p,
		}
		u = ub.String()
		h = http.Header{
			"Authorization": {
				"Basic " + base64.StdEncoding.EncodeToString([]byte(userinfo.String())),
			},
		}
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
