package server

import (
	"net"
	"net/http"

	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
	"k0s.io/pkg/middleware"
)

func init() {
	Tunnels[api.FS] = StartFileServer
	Channels[api.FSID] = StartFileServer
}

func StartFileServer(c agent.Config) chan net.Conn {
	var (
		fsListener = NewLys()
		handler    = middleware.LoggingMiddleware(http.FileServer(http.Dir("/")))
		fileServer = &http.Server{Handler: handler}
	)
	go fileServer.Serve(fsListener)
	return fsListener.Conns
}
