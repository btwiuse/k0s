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

	// NewRPC()
	NewSession() Session

	Close()
	Done() <-chan struct{}

	BasicAuth(http.Handler) http.Handler

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

	// *rpc.Client
	// Call(serviceMethod string, args interface{}, reply interface{}) error
	// Close() error
	// Go(serviceMethod string, args interface{}, reply interface{}, done chan *rpc.Call) *rpc.Call
	NewSession() // Session
}

type RPCManager interface {
	pkg.Manager

	AddRPC(RPC)
	Last() RPC
}
