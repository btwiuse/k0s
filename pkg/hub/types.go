package hub

import (
	"net"
	"net/http"

	// "net/rpc"

	"k0s.io/k0s/pkg"
	"k0s.io/k0s/pkg/api"
	"k0s.io/k0s/pkg/asciitransport"
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
	SessionManager

	// AddRPCConn(net.Conn)
	AddTerminalConn(net.Conn)
	AddSessionConn(net.Conn) // deprecated in favor of AddTerminalConn
	AddMetricsConn(net.Conn)
	AddSocks5Conn(net.Conn)
	AddRedirConn(net.Conn)
	AddFSConn(net.Conn)

	NewTerminal() net.Conn
	NewSession() Session
	NewMetrics() net.Conn
	NewSocks5() net.Conn
	NewRedir() net.Conn
	NewFS() net.Conn

	BasicAuth(http.Handler) http.Handler
}

type SessionManager interface {
	pkg.Manager

	AddSession(Session)
	GetSession(string) Session
}

type TerminalManager interface {
	pkg.Manager

	AddTerminal(Terminal)
	GetTerminal(string) Terminal
}

type Terminal interface {
	pkg.Tider
	asciitransport.AsciiTransportClient
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

	NewTerminal() // Terminal
	NewSession()  // Session
	NewMetrics()
	NewSocks5()
	NewRedir()
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
