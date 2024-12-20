package agent

import (
	"io"
	"net"

	"github.com/btwiuse/version"
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

	GetVersion() *version.Version
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
	TunnelListener
	ChannelChan(api.ProtocolID) chan net.Conn
	AgentRegister(net.Conn) (RPC, error)

	ConnectAndServe() error
	Serve(RPC) error

	// RPC
	// ServeGRPC() error
	// Connect() (RPC, error)
}

type TunnelListener interface {
	AcceptProtocol(api.ProtocolID) (net.Conn, error)
}

type TtyFactory interface {
	MakeTty() (Tty, error)
	MakeTtyCmd([]string) (Tty, error)
	MakeTtyEnv([]string, map[string]string) (Tty, error)
}

type Tty interface {
	io.ReadWriteCloser
	Resize(rows int, cols int) error
}
