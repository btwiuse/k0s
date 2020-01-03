package hub

import (
	"net"
	"net/http"

	// "net/rpc"

	"github.com/btwiuse/conntroll/pkg"
	"github.com/btwiuse/conntroll/pkg/api"
)

type Info interface {
	GetID() string
	GetName() string
	GetTags() []string
	GetAuth() string

	GetOS() string
	GetPwd() string
	GetArch() string
	GetDistro() string
	GetHostname() string
	GetUsername() string
}

type Config interface {
	Port() string
	UseTLS() bool
	LocalUI() bool
	Cert() string
	Key() string
	BasicAuth() (string, string, bool)
}

type Hub interface {
	AgentManager

	// Serve(net.Listener) error
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
	pkg.Tider
	SessionManager

	// AddRPCConn(net.Conn)
	AddSessionConn(net.Conn)

	NewSession() Session

	BasicAuth(http.Handler) http.Handler

	// TODO: remove this
	// currently it is used to generate index.html in localui mode
	GetUsername() string
	GetHostname() string
}

type SessionManager interface {
	pkg.Manager

	AddSession(Session)
	GetSession(string) Session
}

type Session interface {
	pkg.Tider
	api.SessionClient

	// TTY() io.ReaderFrom // | io.WriterTo
	// FS(api.ChunkRequest) io.ReaderFrom
}

type RPC interface {
	pkg.Tider

	Close()
	Done() <-chan struct{}

	NewSession() // Session
	Ping()
	RemoteIP() string
	Actions() <-chan func(Hub)
	Unregister(Hub)
}

type RPCManager interface {
	pkg.Manager

	AddRPC(RPC)
	Last() RPC
}
