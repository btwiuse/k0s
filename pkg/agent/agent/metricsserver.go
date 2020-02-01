package agent

import (
	"net/http"

	types "k0s.io/conntroll/pkg/agent"
	"k0s.io/conntroll/pkg/exporter"
)

func StartMetricsServer(c types.Config) types.MetricsServer {
	var (
		metricsListener = NewLys()
		handler         = LoggingMiddleware(GzipMiddleware(exporter.NewHandler()))
		metricsServer   = &http.Server{Handler: handler}
	)
	go metricsServer.Serve(metricsListener)
	return metricsListener
}
