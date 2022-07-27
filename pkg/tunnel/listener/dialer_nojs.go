//go:build !js

package listener

import (
	"crypto/tls"
	"net/http"

	"nhooyr.io/websocket"
)

var opts = &websocket.DialOptions{
	HTTPClient: &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	},
}
