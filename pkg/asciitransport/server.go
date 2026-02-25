package asciitransport

import (
	"io"
	"sync"
	"time"
)

// ServerSession is the server-side terminal session interface.
// It implements io.ReadWriteCloser, making it usable like a local file or pipe.
type ServerSession interface {
	// io.Reader: reads InputMsg data sent by the client.
	// When an Input message is received, its Data is unpacked and filled into p.
	// Resize messages are queued internally and not returned as io.Reader data.
	io.Reader

	// io.Writer: sends OutputMsg data to the client.
	// Received []byte is automatically packed as [0, "o", string(p)] and sent.
	io.Writer

	// io.Closer: closes the underlying connection.
	io.Closer

	// NextResize blocks until the next window resize event is received.
	// Callers can use a dedicated goroutine to block on this method
	// and adjust the PTY window size accordingly.
	NextResize() (*ResizeMsg, error)
}

type serverSession struct {
	at      *AsciiTransport
	readBuf []byte
}

func (s *serverSession) Read(p []byte) (int, error) {
	if len(s.readBuf) > 0 {
		n := copy(p, s.readBuf)
		s.readBuf = s.readBuf[n:]
		return n, nil
	}
	select {
	case ie := <-s.at.iech:
		data := []byte(ie.Data)
		n := copy(p, data)
		if n < len(data) {
			s.readBuf = data[n:]
		}
		ie.Time = time.Since(s.at.start).Seconds()
		s.at.log(ie)
		return n, nil
	case <-s.at.quit:
		return 0, io.EOF
	}
}

func (s *serverSession) Write(p []byte) (int, error) {
	select {
	case <-s.at.quit:
		return 0, io.ErrClosedPipe
	default:
	}
	s.at.Output(p)
	return len(p), nil
}

func (s *serverSession) Close() error {
	return s.at.Close()
}

func (s *serverSession) NextResize() (*ResizeMsg, error) {
	select {
	case re := <-s.at.rech:
		re.Timestamp = uint(time.Now().Unix())
		s.at.log(re)
		return re, nil
	case <-s.at.quit:
		return nil, io.EOF
	}
}

func Server(conn io.ReadWriteCloser, opts ...Opt) ServerSession {
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
	return &serverSession{at: at}
}
