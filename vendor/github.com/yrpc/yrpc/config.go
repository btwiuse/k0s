package yrpc

import (
	"net"
	"time"

	"github.com/go-kit/kit/metrics"
)

// ServerConfig contains binding infos
type ServerConfig struct {
	Addr                string
	Handler             Handler // handler to invoke
	DefaultReadTimeout  int
	DefaultWriteTimeout int
	ReadFrameChSize     int
	WriteFrameChSize    int
	MaxFrameSize        int
	MaxCloseRate        int // per second
	ListenFunc          func(network, address string) (net.Listener, error)
	OverlayNetwork      func(net.Listener) Listener
	OnKickCB            func(w FrameWriter)
	LatencyMetric       metrics.Histogram
	CounterMetric       metrics.Counter
	ln                  Listener
}

// SubFunc for subscribe callback
type SubFunc func(*Connection, *Frame)

// ClientConfig is conf for Connection
type ClientConfig struct {
	WriteTimeout     int
	ReadTimeout      int
	DialTimeout      time.Duration
	WriteFrameChSize int
	Handler          Handler
	OverlayNetwork   func(address string, dialConfig DialConfig) (net.Conn, error)
}

// DialConfig for dial
type DialConfig struct {
	DialTimeout time.Duration
}
