package server

import (
	"net"
	"net/http"

	"github.com/btwiuse/pretty"
	"k0s.io/pkg/agent/config"
	"k0s.io/pkg/middleware"
)

func StartVersionServer(c *config.Config) chan net.Conn {
	var (
		versionListener = NewChannelListener()
		handler         = middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(pretty.JSONString(c.Version)))
		}))
		versionServer = &http.Server{Handler: handler}
	)
	go versionServer.Serve(versionListener)
	return versionListener.Conns
}
