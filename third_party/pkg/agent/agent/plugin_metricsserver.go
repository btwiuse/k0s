//go:build plugin_metricsserver

package server

import (
	"net"
	"net/http"

	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
	"k0s.io/pkg/middleware"
	"k0s.io/third_party/pkg/exporter"
)

func init() { Tunnels[api.Metrics] = StartMetricsServer }

func StartMetricsServer(c agent.Config) chan net.Conn {
	var (
		metricsListener = NewLys()
		handler         = middleware.LoggingMiddleware(middleware.GzipMiddleware(exporter.NewHandler()))
		metricsServer   = &http.Server{Handler: handler}
	)
	go metricsServer.Serve(metricsListener)
	return metricsListener.Conns
}
