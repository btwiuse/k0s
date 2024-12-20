package api

import "fmt"

type ProtocolID string

const (
	TerminalID   ProtocolID = "TERMINAL"
	SessionID    ProtocolID = "SESSION"
	MetricsID    ProtocolID = "METRICS"
	Socks5ID     ProtocolID = "SOCKS5"
	RedirID      ProtocolID = "REDIR"
	FSID         ProtocolID = "FS"
	PingID       ProtocolID = "PING"
	VersionID    ProtocolID = "VERSION"
	K16sID       ProtocolID = "K16S"
	DohID        ProtocolID = "DOH"
	EnvID        ProtocolID = "ENV"
	TerminalV2ID ProtocolID = "TERMINALV2"
	JsonlID      ProtocolID = "JSONL"
	XpraID       ProtocolID = "XPRA"
)

type Tunnel uint8

const (
	Terminal Tunnel = iota // starts at 0
	Session
	Metrics
	Socks5
	Redir
	FS
	Ping
	Version
	K16s
	Doh
	Env
	TerminalV2
	Jsonl
	Xpra

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
	"K16S",
	"DOH",
	"ENV",
	"TERMINALV2",
	"JSONL",
	"XPRA",
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
	return 0, fmt.Errorf("invalid tunnel type: %s", s)
}
