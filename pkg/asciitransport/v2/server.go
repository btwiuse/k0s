package asciitransport

import (
	"io"
	"sync"
	"time"
)

type AsciiTransportServer interface {
	ResizeEvent() <-chan *Frame_R
	InputEvent() <-chan *Frame_I
	Output([]byte)
	OutputFrom(io.Reader) error
	Done() <-chan struct{}
	Close() error
}

func Server(conn io.ReadWriteCloser, opts ...Opt) AsciiTransportServer {
	at := &AsciiTransport{
		conn:      conn,
		quit:      make(chan struct{}),
		closeonce: &sync.Once{},
		start:     time.Now(),
		iech:      make(chan *Frame_I),
		oech:      make(chan *Frame_O),
		rech:      make(chan *Frame_R),
		isClient:  false,
	}
	for _, opt := range opts {
		opt(at)
	}
	pr, pw := io.Pipe()
	go func() {
		io.Copy(pw, conn)
		at.Close()
	}()
	at.goReadConn(pr)
	at.goWriteConn(conn)
	return at
}
