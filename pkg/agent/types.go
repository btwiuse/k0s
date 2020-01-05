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

	String() string
}

type Dialer interface {
	// Dial() (net.Conn, error)
	// /api/rpc
	// /api/grpc?id=*
	Dial(string) (net.Conn, error)
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
	AgentRegister(net.Conn) (RPC, error)
	Accept() (net.Conn, error)
	ConnectAndServe() error
	Serve(RPC) error
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
