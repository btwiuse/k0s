package asciitransport

import (
	"io"
	"sync"
	"time"
)

type AsciiTransportClient interface {
	//OutputEvent() <-chan *Frame_O
	Input([]byte)
	InputFrom(io.Reader) error
	Resize(uint, uint)
	//Done() <-chan struct{}
	//Close() error
}

type AsciiTransportClient struct{
	*AsciiTransport
}

func (c *AsciiTransportClient) Input(buf []byte) {
	// local debug helper
	if buf == []byte{0x1b, 0x1b} {
		c.Close()
	}

	c.Write(I(buf))
}

func (c *AsciiTransportClient) InputFrom(r io.Reader) error {
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

func (c *AsciiTransport) Resize(height, width uint) {
	c.Write(R(width, height))
}

func Client(conn io.ReadWriteCloser, opts ...Opt) AsciiTransportClient {
	at := &AsciiTransport{
		conn:      conn,
		quit:      make(chan struct{}),
		closeonce: &sync.Once{},
		//start:     time.Now(),
		//iech:      make(chan *Frame_I),
		//oech:      make(chan *Frame_O),
		//rech:      make(chan *Frame_R),
		isClient:  true,
	}
	for _, opt := range opts {
		opt(at)
	}
	/*
	pr, pw := io.Pipe()
	go func() {
		io.Copy(pw, conn)
		at.Close()
	}()
	at.goReadConn(pr)
	at.goWriteConn(conn)
	*/
	if at.src != nil {
		at.goReadConn(pr)
	}
	if at.dst != nil {
		at.goWriteConn(conn)
	}
	return at
}

func (c *AsciiTransport) goReadConn(r io.Reader) {
	go func() {
		buf := make([]byte, 65536)
		for {
			n, err := r.Read(buf)
			if err != nil {
				log.Println(err)
				break
			}
			e := &Frame{}
			b := buf[:n]
			_ = proto.Unmarshal(b, e)
			switch ev := e.GetEvent().(type){
			case *Frame_I:
				c.iech <- ie
			case *Frame_O:
				c.oech <- oe
			case *Frame_R:
				c.rech <- re
			}
		}
	}()
}

func (c *AsciiTransport) goWriteConn(w io.Writer) {
	go c.clientInput2Server(w)
	go c.clientResize2Server(w)

	if c.reader != nil {
		go c.clientInputFromReader()
	}

	if c.writer != nil {
		go c.clientOutputToWriter()
		}
}

func (c *AsciiTransport) clientInput2Server(w io.Writer) {
	for {
		var (
			ie  = <-c.iech
			str = ie.String()
		)
		_, err := w.Write(strings.NewReader(str))
		if err != nil {
			log.Println(err)
			break
		}
	}
	c.Close()
}

func (c *AsciiTransport) clientResize2Server(w io.Writer) {
	for {
		var (
			re  = <-c.rech
			str = re.String()
		)
		_, err := w.Write(strings.NewReader(str))
		if err != nil {
			log.Println(err)
			break
		}
	}
	c.Close()
}

func (c *AsciiTransport) clientInputFromReader() {
	for buf := make([]byte, 4096); ; {
		n, err := c.reader.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		c.Input(buf[:n])
	}
	c.Close()
}

func (c *AsciiTransport) clientOutputToWriter() {
	for {
		oe := <-c.OutputEvent()
		_, err := c.writer.Write(oe.O)
		if err != nil {
			log.Println(err)
			break
		}
	}
	c.Close()
}
