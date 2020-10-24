package asciitransport

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"
)

type Opt func(at *AsciiTransport)

// modeled after io.ReadWriteCloser
interface AsciiTransport {
	Read() (*Frame, error)
	Write(*Frame) error
	Close() error
	Done() <-chan struct{}
}

// Read
func (at *AsciiTransport) Read() (*Frame, error) {
	buf := make([]byte, 65536)
	n, err := r.Read(buf)
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
func (at *AsciiTransport) Write(*Frame) error {
}

type AsciiTransport struct {
	conn      io.ReadWriteCloser
	quit      chan struct{}
	closeonce *sync.Once
	isClient  bool
	src       io.Reader
	dst       io.Writer
}

func (c *AsciiTransport) Close() error {
	c.closeonce.Do(func() {
		close(c.quit)
		c.conn.Close()
		if c.logger != nil {
			c.logger.Close()
		}
	})
	return nil
}

func (c *AsciiTransport) Done() <-chan struct{} {
	return c.quit
}
