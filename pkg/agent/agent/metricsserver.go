package agent

import (
	"net"
	"net/http"

	types "k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/exporter"
)

func StartMetricsServer(c types.Config) chan net.Conn {
	var (
		metricsListener = NewLys()
		handler         = LoggingMiddleware(GzipMiddleware(exporter.NewHandler()))
		metricsServer   = &http.Server{Handler: handler}
	)
	go metricsServer.Serve(metricsListener)
	return metricsListener.Conns
}
