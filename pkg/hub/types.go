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

	ChannelChan(api.ProtocolID) chan net.Conn
	OpenChannel(api.ProtocolID) net.Conn
	BasicAuth(http.Handler) http.Handler
}

type Session interface {
	k0s.Tider
	RemoteIP() string

	Close()
	Done() <-chan struct{}

	OpenChannel(api.ProtocolID)

	Ping()
	Actions() <-chan func(Hub)
}
