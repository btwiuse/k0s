//go:build !gorilla
// +build !gorilla

package dial

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"net"
	"net/http"
	"net/url"

	"nhooyr.io/websocket"
)

func Dial(u *url.URL) (conn net.Conn, err error) {
	wsconn, _, err := websocket.Dial(
		context.Background(),
		u.String(),
		&websocket.DialOptions{
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyFromEnvironment,
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			},
			HTTPHeader: authorizationHeader(u.User),
		},
	)
	if err != nil {
		return nil, err
	}

	return websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}

func authorizationHeader(userinfo *url.Userinfo) http.Header {
	return http.Header{
		"Authorization": {
			"Basic " + base64.StdEncoding.EncodeToString([]byte(userinfo.String())),
		},
	}
}
