package agent

import (
	"io"
	"net"

	"k0s.io/conntroll/pkg"
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
	GetReadOnly() bool
	GetVerbose() bool
	GetInsecure() bool

	String() string

	GetVersion() pkg.Version
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
	AcceptGrpc() (net.Conn, error)
	AcceptSocks5() (net.Conn, error)
	ConnectAndServe() error
	Serve(RPC) error
	GrpcChanConn() chan<- net.Conn
	Socks5ChanConn() chan<- net.Conn
	// RPC
	// ServeGRPC() error
	// Connect() (RPC, error)
}

type Socks5Server interface {
	ChanConn() chan<- net.Conn
}

type GrpcServer interface {
	ChanConn() chan<- net.Conn
}

type TtyFactory interface {
	MakeTty() (Tty, error)
}

type Tty interface {
	io.ReadWriteCloser
	Resize(rows int, cols int) error
}
