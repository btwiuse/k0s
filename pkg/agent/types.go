package agent

import (
	"io"
	"net"
)

type Config interface {
	ID() string
	Name() string
	Tags() []string
	Port() string
	Addr() string
	Scheme() string
	Verbose() bool
	Insecure() bool
	Hostname() string
	NewAgentRequestBody() []byte
	NewSessionRequestBody() []byte
}

type Agent interface {
	TtyFactory

	CreateSession() (net.Conn, error)
	ConnectAndServe() error
	Connect() (net.Conn, error)
	Dial() (net.Conn, error)

	Go(func() error)
	Wait() error
}

type TtyFactory interface {
	MakeTty() (Tty, error)
}

type Tty interface {
	io.ReadWriteCloser
	Resize(rows int, cols int) error
}
