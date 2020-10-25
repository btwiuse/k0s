package asciitransport

import (
	"io"
	"log"
	"sync"

	proto "github.com/golang/protobuf/proto"
)

type Opt func(at *AsciiTransport)

// modeled after io.ReadWriteCloser
/*
interface AsciiTransport {
	Read() (*Frame, error)
	Write(*Frame) error
	Close() error
	Done() <-chan struct{}
}
*/

// Read
// Unmarshal
func (at *AsciiTransport) Read() (*Frame, error) {
	buf := make([]byte, 65536)
	n, err := at.conn.Read(buf)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	e := &Frame{}
	b := buf[:n]
	err = proto.Unmarshal(b, e)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return e, err
}

// Write
// Marshal
func (at *AsciiTransport) Write(f *Frame) error {
	b, _ := proto.Marshal(f)
	_, err := at.conn.Write(b)
	if err != nil {
		return err
	}
	return nil
}

type AsciiTransport struct {
	conn       io.ReadWriteCloser
	quit       chan struct{}
	closeonce  *sync.Once
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
