package agent

import (
	"io"
	"net"

	"k0s.io/pkg/agent/config"
	"k0s.io/pkg/api"
)

type Session interface {
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
	Config() *config.Config
	TunnelListener
	ChannelChan(api.ProtocolID) chan net.Conn
	AgentRegister(net.Conn) (Session, error)

	ConnectAndServe() error
	Serve(Session) error
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
