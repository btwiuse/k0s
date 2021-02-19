package agent

import (
	"io"
	"net"

	"k0s.io/pkg"
	"k0s.io/pkg/api"
)

type Info interface {
	GetOS() string
	GetPwd() string
	GetArch() string
	GetDistro() string
	GetHostname() string
	GetUsername() string
}

type Config interface {
	Info

	GetID() string
	GetName() string
	GetTags() []string

	GetHost() string
	GetPort() string
	GetAddr() string
	GetScheme() string
	GetSchemeWS() string

	GetCmd() []string
	GetReadOnly() bool
	GetVerbose() bool
	GetInsecure() bool
	GetPet() bool

	String() string

	GetVersion() pkg.Version
}

type Dialer interface {
	// Dial() (net.Conn, error)
	// /api/rpc
	// /api/grpc?id=*
	Dial(string, string) (net.Conn, error)
}

type RPC interface {
	// NewConnection()
	// Ping()
	// Pong()
	// Connect() (net.Conn, error)
	Close()
	Done() <-chan struct{}
	Actions() <-chan func(Agent)
}

// inside agent there are:
// config (static)
// grpcServer (long lived)
// rpc client/server (ephemeral)
type Agent interface {
	Config
	Dialer
	TunnelListener
	TunnelChan(api.Tunnel) chan net.Conn
	AgentRegister(net.Conn) (RPC, error)

	ConnectAndServe() error
	Serve(RPC) error

	// RPC
	// ServeGRPC() error
	// Connect() (RPC, error)
}

type TunnelListener interface {
	Accept(api.Tunnel) (net.Conn, error)
}

type TtyFactory interface {
	MakeTty() (Tty, error)
}

type Tty interface {
	io.ReadWriteCloser
	Resize(rows int, cols int) error
}
