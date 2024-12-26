package api

type ProtocolID = string

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
