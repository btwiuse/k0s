package hub

import (
	"net"
	"net/http"

	// "net/rpc"

	"k0s.io"
	"k0s.io/pkg/api"
)

type AgentInfo interface {
	GetID() string
	GetName() string
	GetTags() []string
	GetAuth() bool
	GetHtpasswd() map[string]string

	GetOS() string
	GetPwd() string
	GetArch() string
	GetDistro() string
	GetHostname() string
	GetUsername() string
	GetVersion() string
	GetGitSummary() string

	SetIP(string)
}

type Config interface {
	Port() string
	UseTLS() bool
	UI() bool
	Verbose() bool
	Cert() string
	Key() string
	GetVersion() k0s.Version
}

type Hub interface {
	AgentManager

	// Serve(net.Listener) error
	GetConfig() Config
	ListenAndServe() error
	ListenAndServeTLS(certFile, keyFile string) error
	Serve(ln net.Listener) error
	ServeTLS(ln net.Listener, certFile, keyFile string) error
}

type AgentManager interface {
	k0s.Manager

	AddAgent(Agent)
	GetAgent(string) Agent
	GetAgents() []Agent
}

type Agent interface {
	AgentInfo
	k0s.Tider

	AddTunnel(api.Tunnel, net.Conn)
	NewTunnel(api.Tunnel) net.Conn
	BasicAuth(http.Handler) http.Handler
}

type RPC interface {
	k0s.Tider

	Close()
	Done() <-chan struct{}

	NewTunnel(api.Tunnel)

	Ping()
	RemoteIP() string
	Actions() <-chan func(Hub)
}
