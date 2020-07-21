package hub

import (
	"net"
	"net/http"

	// "net/rpc"

	"k0s.io/k0s/pkg"
	"k0s.io/k0s/pkg/api"
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

	SetIP(string)
}

type Config interface {
	Port() string
	UseTLS() bool
	Cert() string
	Key() string
	GetVersion() pkg.Version
}

type Hub interface {
	AgentManager

	// Serve(net.Listener) error
	GetConfig() Config
	ListenAndServe() error
	ListenAndServeTLS(certFile, keyFile string) error
}

type AgentManager interface {
	pkg.Manager

	AddAgent(Agent)
	GetAgent(string) Agent
	GetAgents() []Agent
}

type Agent interface {
	AgentInfo
	pkg.Tider

	AddTunnel(api.Tunnel, net.Conn)
	NewTunnel(api.Tunnel) net.Conn
	BasicAuth(http.Handler) http.Handler
}

type RPC interface {
	pkg.Tider

	Close()
	Done() <-chan struct{}

	NewTunnel(api.Tunnel)

	Ping()
	RemoteIP() string
	Actions() <-chan func(Hub)
}
