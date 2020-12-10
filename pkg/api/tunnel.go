package api

import "fmt"

type Tunnel uint8

//go:generate stringer -type=Tunnel

const (
	Terminal Tunnel = iota // starts at 0
	Session
	Metrics
	Socks5
	Redir
	FS
	Ping
	Version

	MaxTunnel // number of tunnels
)

var tunnelNames = []string{
	"TERMINAL",
	"SESSION",
	"METRICS",
	"SOCKS5",
	"REDIR",
	"FS",
	"PING",
	"VERSION",
}

func (tun Tunnel) String() string {
	return tunnelNames[tun]
}

func FromString(s string) (Tunnel, error) {
	for i := range tunnelNames {
		if s == tunnelNames[i] {
			return Tunnel(i), nil
		}
	}
	return 0, fmt.Errorf("Invalid tunnel type: %s", s)
}
