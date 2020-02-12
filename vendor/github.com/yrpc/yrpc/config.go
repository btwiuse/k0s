package yrpc

import (
	"net"

	"github.com/go-kit/kit/metrics"
)

// ServerConfig contains binding infos
type ServerConfig struct {
	Addr             string
	Handler          Handler // handler to invoke
	ReadFrameChSize  int
	WriteFrameChSize int
	MaxCloseRate     int // per second
	Listener         net.Listener
	OnKickCB         func(w FrameWriter)
	LatencyMetric    metrics.Histogram
	CounterMetric    metrics.Counter
}

// SubFunc for subscribe callback
type SubFunc func(*Connection, *Frame)

// ClientConfig is conf for Connection
type ClientConfig struct {
	WriteFrameChSize int
	Handler          Handler
	Dialer           func() (net.Conn, error)
}
