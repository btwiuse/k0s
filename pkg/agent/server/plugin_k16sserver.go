//go:build plugin_k16sserver
// +build plugin_k16sserver

package server

import (
	"net"
	"net/http"

	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
	"k0s.io/pkg/exporter/k16s"
	"k0s.io/pkg/middleware"
)

func init() { Tunnels[api.K16s] = StartK16sServer }

func StartK16sServer(c agent.Config) chan net.Conn {
	var (
		k16sListener = NewLys()
		handler      = middleware.LoggingMiddleware(middleware.GzipMiddleware(k16s.NewHandler()))
		k16sServer   = &http.Server{Handler: handler}
	)
	go k16sServer.Serve(k16sListener)
	return k16sListener.Conns
}
