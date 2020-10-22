package asciitransport

import (
	"io"
	"sync"
	"time"
)

type AsciiTransportClient interface {
	OutputEvent() <-chan *Frame_O
	Input([]byte)
	InputFrom(io.Reader) error
	Resize(uint, uint)
	Done() <-chan struct{}
	Close() error
}

func Client(conn io.ReadWriteCloser, opts ...Opt) AsciiTransportClient {
	at := &AsciiTransport{
		conn:      conn,
		quit:      make(chan struct{}),
		closeonce: &sync.Once{},
		start:     time.Now(),
		iech:      make(chan *Frame_I),
		oech:      make(chan *Frame_O),
		rech:      make(chan *Frame_R),
		isClient:  true,
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
