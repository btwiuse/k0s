package agent

type Config interface {
	ID() string
	Port() string
	Addr() string
	Scheme() string
	Hostname() string
}

type Agent interface {
	ID() string
	OS() string
	Pwd() string
	ARCH() string
	Username() string
	Hostname() string

	// MkTTYSession
	// Session
}

type Client interface {
	ListAgents() []Agent
}
