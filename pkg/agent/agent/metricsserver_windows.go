// +build windows
// +build !linux
// +build !darwin

package agent

import (
	"net/http"

	types "k0s.io/conntroll/pkg/agent"
)

func StartMetricsServer(c types.Config) types.MetricsServer {
	var (
		metricsListener = NewLys()
		handler         = LoggingMiddleware(GzipMiddleware(http.NewServeMux()))
		metricsServer   = &http.Server{Handler: handler}
	)
	go metricsServer.Serve(metricsListener)
	return metricsListener
}
