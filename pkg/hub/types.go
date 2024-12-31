package hub

import (
	"net"
	"net/http"

	// "net/rpc"

	"k0s.io"
	"k0s.io/pkg/api"
	"k0s.io/pkg/hub/agent/info"
	"k0s.io/pkg/hub/config"
)

type Hub interface {
	AgentManager

	Config() *config.Config
	Handler() http.Handler
}

type AgentManager interface {
	k0s.Manager

	AddAgent(Agent)
	GetAgent(string) Agent
	GetAgents() []Agent
}

type Agent interface {
	Info() *info.Info
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
