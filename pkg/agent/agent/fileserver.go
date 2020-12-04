package agent

import (
	"net"
	"net/http"

	types "k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/middleware"
)

func StartFileServer(c types.Config) chan net.Conn {
	var (
		fsListener = NewLys()
		handler    = middleware.LoggingMiddleware(http.FileServer(http.Dir("/")))
		fileServer = &http.Server{Handler: handler}
	)
	go fileServer.Serve(fsListener)
	return fsListener.Conns
}
