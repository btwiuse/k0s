package agent

import (
	"io"
	"net"
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

	GetCmd() []string
	GetVerbose() bool
	GetInsecure() bool

	FakeHeader(p string) []byte

	String() string
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
	AgentRegister(net.Conn) (RPC, error)
	Accept() (net.Conn, error)
	Dial() (net.Conn, error)
	ConnectAndServe() error
	Serve(RPC)
	GRPCServer
	// RPC
	// ServeGRPC() error
	// Connect() (RPC, error)
}

type GRPCServer interface {
	ChanConn() chan<- net.Conn
	// TtyFactory
}

type TtyFactory interface {
	MakeTty() (Tty, error)
}

type Tty interface {
	io.ReadWriteCloser
	Resize(rows int, cols int) error
}
