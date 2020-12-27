package agent

import (
	"net"
	"net/http"

	types "k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/exporter/k16s"
	"k0s.io/k0s/pkg/middleware"
)

func StartK16sServer(c types.Config) chan net.Conn {
	var (
		k16sListener = NewLys()
		handler      = middleware.LoggingMiddleware(middleware.GzipMiddleware(k16s.NewHandler()))
		k16sServer   = &http.Server{Handler: handler}
	)
	go k16sServer.Serve(k16sListener)
	return k16sListener.Conns
}
