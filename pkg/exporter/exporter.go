package exporter

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/prometheus/node_exporter/collector"
	"gopkg.in/alecthomas/kingpin.v2"
)

// handler wraps an unfiltered http.Handler but uses a filtered handler,
// created on the fly, if filtering is requested. Create instances with
// newHandler.
type handler struct {
	innerHandler http.Handler
	// exporterMetricsRegistry is a separate registry for the metrics about
	// the exporter itself.
	exporterMetricsRegistry *prometheus.Registry
}

type Logger struct{}

func (l *Logger) Log(_ ...interface{}) error {
	return nil
}

var logger = &Logger{}

func NewHandler() http.Handler {
	kingpin.MustParse(kingpin.CommandLine.Parse([]string{}))
	h := &handler{
		exporterMetricsRegistry: prometheus.NewRegistry(),
	}
	h.exporterMetricsRegistry.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
	)
	err := h.innerHandlerInit()
	if err != nil {
		panic(fmt.Sprintf("Couldn't create metrics handler: %s", err))
	}
	return h
}

// ServeHTTP implements http.Handler.
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.innerHandler.ServeHTTP(w, r)
	return
}

// innerHandler is used to create both the one unfiltered http.Handler to be
// wrapped by the outer handler and also the filtered handlers created on the
// fly. The former is accomplished by calling innerHandler without any arguments
// (in which case it will log all the collectors enabled via command-line
// flags).
func (h *handler) innerHandlerInit() error {
	nc, err := collector.NewNodeCollector(logger)
	if err != nil {
		return fmt.Errorf("couldn't create collector: %s", err)
	}

	r := prometheus.NewRegistry()
	r.MustRegister(version.NewCollector("node_exporter"))
	if err := r.Register(nc); err != nil {
		return fmt.Errorf("couldn't register node collector: %s", err)
	}
	handler := promhttp.HandlerFor(
		prometheus.Gatherers{h.exporterMetricsRegistry, r},
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
			Registry:      h.exporterMetricsRegistry,
		},
	)
	// Note that we have to use h.exporterMetricsRegistry here to
	// use the same promhttp metrics for all expositions.
	h.innerHandler = promhttp.InstrumentMetricHandler(
		h.exporterMetricsRegistry, handler,
	)
	return nil
}
