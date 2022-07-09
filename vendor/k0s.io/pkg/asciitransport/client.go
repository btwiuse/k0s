package asciitransport

import (
	"io"
	"sync"
	"time"
)

type AsciiTransportClient interface {
	OutputEvent() <-chan *OutputEvent
	Input([]byte)
	InputFrom(io.Reader) error
	Resize(uint, uint)
	Done() <-chan struct{}
	Close() error
}

func Client(conn io.ReadWriteCloser, opts ...Opt) AsciiTransportClient {
	at := &AsciiTransport{
		conn:       conn,
		quit:       make(chan struct{}),
		closeonce:  &sync.Once{},
		start:      time.Now(),
		iech:       make(chan *InputEvent),
		oech:       make(chan *OutputEvent),
		rech:       make(chan *ResizeEvent),
		isClient:   true,
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
