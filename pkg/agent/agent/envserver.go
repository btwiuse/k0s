package agent

import (
	"net"
	"net/http"

	types "k0s.io/pkg/agent"
	"k0s.io/pkg/exporter/env"
	"k0s.io/pkg/middleware"
)

func StartEnvServer(c types.Config) chan net.Conn {
	var (
		envListener = NewLys()
		handler     = middleware.LoggingMiddleware(middleware.GzipMiddleware(env.NewHandler()))
		envServer   = &http.Server{Handler: handler}
	)
	go envServer.Serve(envListener)
	return envListener.Conns
}
