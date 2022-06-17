package asciitransport

import (
	"io"
	"sync"
	"time"
)

type AsciiTransportServer interface {
	ResizeEvent() <-chan *ResizeEvent
	InputEvent() <-chan *InputEvent
	Output([]byte)
	OutputFrom(io.Reader) error
	Done() <-chan struct{}
	Close() error
	ApplyOpts(...Opt)
}

func Server(conn io.ReadWriteCloser, opts ...Opt) AsciiTransportServer {
	at := &AsciiTransport{
		conn:       conn,
		quit:       make(chan struct{}),
		closeonce:  &sync.Once{},
		start:      time.Now(),
		iech:       make(chan *InputEvent),
		oech:       make(chan *OutputEvent),
		rech:       make(chan *ResizeEvent),
		isClient:   false,
		readerOnce: &sync.Once{},
		writerOnce: &sync.Once{},
	}
	at.ApplyOpts(opts...)
	pr, pw := io.Pipe()
	go func() {
		io.Copy(pw, conn)
		at.Close()
	}()
	at.goReadConn(pr)
	at.goWriteConn(conn)
	return at
}
