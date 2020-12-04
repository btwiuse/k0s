package agent

import (
	"net"
	"net/http"

	types "k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/exporter"
	"k0s.io/k0s/pkg/middleware"
)

func StartMetricsServer(c types.Config) chan net.Conn {
	var (
		metricsListener = NewLys()
		handler         = middleware.LoggingMiddleware(middleware.GzipMiddleware(exporter.NewHandler()))
		metricsServer   = &http.Server{Handler: handler}
	)
	go metricsServer.Serve(metricsListener)
	return metricsListener.Conns
}
