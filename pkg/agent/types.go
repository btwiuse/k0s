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

	// NewYRPCAgentRequestBody() []byte
	// NewAgentRequestBody() []byte
	// NewSessionRequestBody() []byte
}

type Agent interface {
	Config

	// TtyFactory

	// create a session capable of either shell or file system access
	// CreateSession() (net.Conn, error)
	Accept() (net.Conn, error)
	// ConnectAndServe() error
	YRPCConnectAndServe() error
	// Connect() (net.Conn, error)
	// Dial() (net.Conn, error)

	// Go(func() error)
	// Wait() error
}

type TtyFactory interface {
	MakeTty() (Tty, error)
}

type Tty interface {
	io.ReadWriteCloser
	Resize(rows int, cols int) error
}
