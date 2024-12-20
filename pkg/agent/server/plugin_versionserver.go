package server

import (
	"net"
	"net/http"

	"github.com/btwiuse/pretty"
	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
	"k0s.io/pkg/middleware"
)

func init() {
	Channels[api.VersionID] = StartVersionServer
}

func StartVersionServer(c agent.Config) chan net.Conn {
	var (
		versionListener = NewLys()
		handler         = middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(pretty.JSONString(c.GetVersion())))
		}))
		versionServer = &http.Server{Handler: handler}
	)
	go versionServer.Serve(versionListener)
	return versionListener.Conns
}
