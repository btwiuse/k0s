package hub

import (
	"net"
	"net/http"

	// "net/rpc"

	"github.com/btwiuse/version"
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
	Ufo() string
	GetVersion() *version.Version
}

type Hub interface {
	AgentManager

	GetConfig() Config
	Handler() http.Handler
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

	AddChannel(api.ProtocolID, net.Conn)
	NewChannel(api.ProtocolID) net.Conn
	BasicAuth(http.Handler) http.Handler
}

type RPC interface {
	k0s.Tider

	Close()
	Done() <-chan struct{}

	NewChannel(api.ProtocolID)

	Ping()
	RemoteIP() string
	Actions() <-chan func(Hub)
}
