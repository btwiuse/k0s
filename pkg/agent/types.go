package agent

import (
	"io"
	"net"
)

type Agent interface {
	TtyFactory

	ConnectAndServe() error
	CreateSession() (net.Conn, error)
	Connect() (net.Conn, error)

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
