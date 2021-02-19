package listener

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"

	portless "k0s.io/pkg/tunnel"

	"nhooyr.io/websocket"
)

func wscheme(u *url.URL) string {
	switch u.Scheme {
	case "https":
		return "wss"
	default:
		return "ws"
	}
}

func Dial(addr string, from string) (conn net.Conn, err error) {
	_url, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	var (
		ub = &url.URL{
			Scheme:   wscheme(_url),
			Host:     _url.Host,
			Path:     _url.Path,
			RawQuery: "from=" + from,
		}
		u = ub.String()
		t = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		opts = &websocket.DialOptions{
			HTTPClient: t,
			HTTPHeader: make(map[string][]string),
		}
	)
	opts.HTTPHeader.Set(portless.FingerPrintHeader, fingerprint)

	wsconn, _, err := websocket.Dial(context.Background(), u, opts)
	if err != nil {
		return nil, err
	}

	return websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}
