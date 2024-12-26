package api

type ProtocolID = string

const (
	TerminalID   ProtocolID = "terminal"
	SessionID    ProtocolID = "session"
	MetricsID    ProtocolID = "metrics"
	Socks5ID     ProtocolID = "socks5"
	RedirID      ProtocolID = "redir"
	FSID         ProtocolID = "fs"
	PingID       ProtocolID = "ping"
	VersionID    ProtocolID = "version"
	K16sID       ProtocolID = "k16s"
	DohID        ProtocolID = "doh"
	EnvID        ProtocolID = "env"
	TerminalV2ID ProtocolID = "terminalv2"
	JsonlID      ProtocolID = "jsonl"
	XpraID       ProtocolID = "xpra"
)
