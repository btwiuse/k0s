package listener

import (
	"context"
	"net"
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

	ub := &url.URL{
		Scheme: wscheme(_url),
		Host:   _url.Host,
		Path:   _url.Path,
	}

	values := ub.Query()
	values.Set(portless.FingerPrintHeader, fingerprint)
	values.Set("from", from)
	ub.RawQuery = values.Encode()

	u := ub.String()
	println(u)

	wsconn, _, err := websocket.Dial(context.Background(), u, opts)
	if err != nil {
		return nil, err
	}

	return websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}
