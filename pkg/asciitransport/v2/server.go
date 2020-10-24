package asciitransport

import (
	"io"
	"sync"
	"time"
)

type AsciiTransportServer interface {
	//ResizeEvent() <-chan *Frame_R
	//InputEvent() <-chan *Frame_I
	Output([]byte)
	OutputFrom(io.Reader) error
	//Done() <-chan struct{}
	//Close() error
}

type AsciiTransportServer struct{
	*AsciiTransport
}

func (c *AsciiTransportServer) Output(buf []byte) {
	c.Write(O(buf))
}

func (c *AsciiTransportServer) OutputFrom(r io.Reader) error {
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

func Server(conn io.ReadWriteCloser, opts ...Opt) AsciiTransportServer {
	at := &AsciiTransport{
		conn:      conn,
		quit:      make(chan struct{}),
		closeonce: &sync.Once{},
		start:     time.Now(),
		//iech:      make(chan *Frame_I),
		//oech:      make(chan *Frame_O),
		//rech:      make(chan *Frame_R),
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

func (c *AsciiTransport) goWriteConn(w io.Writer) {
	go c.serverOutput2Client(w)
	go c.serverOutputPing2Client(w)

	if c.reader != nil {
		go c.serverOutputFromReader()
	}

	if c.writer != nil {
		go c.serverInputToWriter()
	}
}


func (c *AsciiTransport) serverOutput2Client(w io.Writer) {
	for {
		var (
			oe  = <-c.oech
			str = oe.String()
		)
		_, err := io.Copy(w, strings.NewReader(str))
		if err != nil {
			log.Println(err)
			break
		}
		oe.Time = time.Since(c.start).Seconds()
		c.log(oe)
	}
	c.Close()
}

func (c *AsciiTransport) serverOutputPing2Client(w io.Writer) {
	for {
		var (
			pe  = (*PingEvent)(&Event{Time: 0, Type: "o", Data: ""})
			str = pe.String()
		)
		_, err := io.Copy(w, strings.NewReader(str))
		if err != nil {
			log.Println(err)
			break
		}
		pe.Time = time.Since(c.start).Seconds()
		c.log(pe)
		time.Sleep(5 * time.Second)
	}
	c.Close()
}

func (c *AsciiTransport) serverOutputFromReader() {
	for buf := make([]byte, 4096); ; {
		n, err := c.reader.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		c.Output(buf[:n])
	}
	c.Close()
}

func (c *AsciiTransport) serverInputToWriter() {
	for {
		ie := <-c.InputEvent()
		_, err := io.Copy(c.writer, strings.NewReader(ie.Data))
		if err != nil {
			log.Println(err)
			break
		}
	}
	c.Close()
}

