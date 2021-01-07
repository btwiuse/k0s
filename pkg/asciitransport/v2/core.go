package asciitransport

import (
	"io"
	"log"
	"sync"

	proto "github.com/golang/protobuf/proto"

	"k0s.io/k0s/pkg/asciiproto"
)

// modeled after io.ReadWriteCloser
/*
interface AsciiTransport {
	Read() (*asciiproto.Frame, error)
	Write(*asciiproto.Frame) error
	Close() error
	Done() <-chan struct{}
}
*/

// Read
// Unmarshal
func (at *AsciiTransport) Read() (*asciiproto.Frame, error) {
	buf := make([]byte, 5536)
	n, err := at.conn.Read(buf)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	f := &asciiproto.Frame{}
	b := buf[:n]
	err = proto.Unmarshal(b, f)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	at.log(f)
	return f, err
}

// Write
// Marshal
func (at *AsciiTransport) Write(f *asciiproto.Frame) error {
	b, _ := proto.Marshal(f)
	_, err := at.conn.Write(b)
	if err != nil {
		return err
	}
	at.log(f)
	return nil
}

type AsciiTransport struct {
	conn       io.ReadWriteCloser
	quit       chan struct{}
	closeonce  *sync.Once
	logger     Logger
	isClient   bool
	src        io.Reader
	dst        io.Writer
	resizehook func(_, _ uint16)
}

func (c *AsciiTransport) Close() error {
	c.closeonce.Do(func() {
		close(c.quit)
		c.conn.Close()
	})
	return nil
}

func (c *AsciiTransport) Done() <-chan struct{} {
	return c.quit
}

func (c *AsciiTransport) log(f *asciiproto.Frame) {
	if c.logger != nil {
		c.logger.Log(f)
	}
}
