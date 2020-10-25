package asciitransport

import (
	"io"
	"sync"
)

/*
type AsciiTransportClient interface {
	//OutputEvent() <-chan *Frame_O
	Input([]byte)
	InputFrom(io.Reader) error
	Resize(uint, uint)
	//Done() <-chan struct{}
	//Close() error
}
*/

type AsciiTransportClient struct {
	*AsciiTransport
}

func (c *AsciiTransportClient) Input(buf []byte) {
	// local debug helper
	// if buf == []byte{0x1b, 0x1b} {
	if string(buf) == "0x1b0x1b" {
		c.Close()
	}

	c.Write(I(buf))
}

func (c *AsciiTransportClient) InputFrom(r io.Reader) error {
	defer c.Close()
	// make([]byte, 0, 4096) causes 0 return
	for buf := make([]byte, 4096); ; {
		n, err := r.Read(buf)
		if err != nil {
			return err
		}
		c.Input(buf[:n])
	}
	return nil
}

func (c *AsciiTransportClient) Resize(height, width uint16) {
	c.Write(R(width, height))
}

func (c *AsciiTransportClient) ReadLoop() {
	defer c.Close()
	for {
		frame, err := c.Read()
		if err != nil {
			break
		}
		buf := frame.Event.(*Frame_O).O
		if len(buf) == 0 || c.dst == nil {
			continue
		}
		c.dst.Write(buf)
	}
}

func Client(conn io.ReadWriteCloser, opts ...Opt) *AsciiTransportClient {
	atc := &AsciiTransportClient{
		AsciiTransport: &AsciiTransport{
			conn:      conn,
			quit:      make(chan struct{}),
			closeonce: &sync.Once{},
			//start:     time.Now(),
			//iech:      make(chan *Frame_I),
			//oech:      make(chan *Frame_O),
			//rech:      make(chan *Frame_R),
			isClient: true,
		},
	}
	for _, opt := range opts {
		opt(atc.AsciiTransport)
	}
	if atc.src != nil {
		go atc.InputFrom(atc.src)
	}
	go atc.ReadLoop()
	return atc
}
