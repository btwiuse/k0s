//go:build !js

package listener

import (
	"crypto/tls"
	"net/http"

	"github.com/coder/websocket"
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
