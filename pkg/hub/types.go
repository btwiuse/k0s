package hub

import (
	"net"
	"net/http"

	// "net/rpc"

	"k0s.io/conntroll/pkg"
	"k0s.io/conntroll/pkg/api"
)

type Info interface {
	GetID() string
	GetName() string
	GetTags() []string
	GetHtpasswd() map[string]string

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
	pkg.Tider
	SessionManager

	// AddRPCConn(net.Conn)
	AddSessionConn(net.Conn)
	AddSocks5Conn(net.Conn)
	AddFSConn(net.Conn)

	NewSession() Session
	NewSocks5() net.Conn
	NewFS() net.Conn

	BasicAuth(http.Handler) http.Handler
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

	// close the underlying *grpc.ClientConn to release memory
	Close() error
}

type RPC interface {
	pkg.Tider

	Close()
	Done() <-chan struct{}

	NewSession() // Session
	NewSocks5()
	NewFS()

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
