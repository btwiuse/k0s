package asciitransport

import (
	"io"
	"sync"
	"time"

	"k0s.io/pkg/asciiproto"
)

/*
type AsciiTransportServer interface {
	//ResizeEvent() <-chan *asciiproto.Frame_R
	//InputEvent() <-chan *asciiproto.Frame_I
	Output([]byte)
	OutputFrom(io.Reader) error
	//Done() <-chan struct{}
	//Close() error
}
*/

type AsciiTransportServer struct {
	*AsciiTransport
}

func (c *AsciiTransportServer) Output(buf []byte) {
	c.Write(O(buf))
}

func (c *AsciiTransportServer) OutputFrom(r io.Reader) error {
	defer c.Close()
	// make([]byte, 0, 4096) causes 0 return
	for buf := make([]byte, 4096); ; {
		n, err := r.Read(buf)
		if err != nil {
			return err
		}
		c.Output(buf[:n])
	}
	return nil
}

func (c *AsciiTransportServer) ReadLoop() {
	defer c.Close()
	for {
		frame, err := c.Read()
		if err != nil {
			break
		}
		switch e := frame.Event.(type) {
		case *asciiproto.Frame_I:
			buf := e.I
			if len(buf) == 0 || c.dst == nil {
				continue
			}
			c.dst.Write(buf)
		case *asciiproto.Frame_R:
			if c.resizehook == nil {
				continue
			}
			u := e.R
			w, h := Uw(u), Uh(u)
			c.resizehook(w, h)
		}
	}
}

func Server(conn io.ReadWriteCloser, opts ...Opt) *AsciiTransportServer {
	ats := &AsciiTransportServer{
		AsciiTransport: &AsciiTransport{
			conn:      conn,
			quit:      make(chan struct{}),
			closeonce: &sync.Once{},
			isClient:  false,
		},
	}
	for _, opt := range opts {
		opt(ats.AsciiTransport)
	}
	if ats.AsciiTransport.src != nil {
		go ats.OutputFrom(ats.AsciiTransport.src)
	}
	go ats.ReadLoop()
	go ats.serverPing()
	return ats
}

func (c *AsciiTransportServer) serverPing() {
	defer c.Close()
	for {
		c.Write(O([]byte{}))
		time.Sleep(5 * time.Second)
	}
}
