package server

import (
	"net"
	"net/http"

	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
	"k0s.io/pkg/middleware"
	"k0s.io/third_party/pkg/exporter/env"
)

func init() { Tunnels[api.Env] = StartEnvServer }

func StartEnvServer(c agent.Config) chan net.Conn {
	var (
		envListener = NewLys()
		handler     = middleware.LoggingMiddleware(middleware.GzipMiddleware(env.NewHandler()))
		envServer   = &http.Server{Handler: handler}
	)
	go envServer.Serve(envListener)
	return envListener.Conns
}
