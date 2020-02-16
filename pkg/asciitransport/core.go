package asciitransport

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"
)

type Opt func(at *AsciiTransport)

type AsciiTransport struct {
	conn      io.ReadWriteCloser
	quit      chan struct{}
	closeonce *sync.Once
	start     time.Time
	logger    Logger
	iech      chan *InputEvent
	oech      chan *OutputEvent
	rech      chan *ResizeEvent
	isClient  bool
	reader    io.Reader
	writer    io.Writer
	resizer   Resizer
}

func (c *AsciiTransport) OutputEvent() <-chan *OutputEvent { return c.oech }
func (s *AsciiTransport) InputEvent() <-chan *InputEvent   { return s.iech }
func (s *AsciiTransport) ResizeEvent() <-chan *ResizeEvent { return s.rech }

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

func (c *AsciiTransport) Input(buf []byte) {
	var (
		str = string(buf)
		e   = &Event{Time: 0, Type: "i", Data: str}
		ie  = (*InputEvent)(e)
	)

	// local debug helper
	if fmt.Sprintf("%q", str) == `"\x1b\x1b"` {
		c.Close()
	}

	c.iech <- ie
}

func (c *AsciiTransport) InputFrom(r io.Reader) error {
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

func (c *AsciiTransport) OutputFrom(r io.Reader) error {
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

func (c *AsciiTransport) Output(buf []byte) {
	var (
		str = string(buf)
		e   = &Event{Time: 0, Type: "o", Data: str}
		oe  = (*OutputEvent)(e)
	)

	c.oech <- oe
}

func (c *AsciiTransport) Resize(height, width uint) {
	ie := &ResizeEvent{
		Version: 2,
		Width:   width,
		Height:  height,
	}
	c.rech <- ie
}

func (c *AsciiTransport) log(v interface{}) {
	if c.logger != nil {
		c.logger.Print(v)
	}
}

func (c *AsciiTransport) goReadConn(r io.Reader) {
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			var (
				buf  = scanner.Bytes()
				line = scanner.Text()
			)
			if len(line) == 0 {
				continue
			}
			if line[0] == '[' {
				var (
					e   = new(Event)
					err = json.Unmarshal(buf, e)
				)
				if err != nil {
					log.Println(err)
					continue
				}
				switch e.Type {
				case "i":
					var (
						ie = (*InputEvent)(e)
						// str = ie.String()
					)
					// consumed by reading <-AsciiTransportServer.OutputEvent()
					c.iech <- ie
					ie.Time = time.Since(c.start).Seconds()
					c.log(ie)
				case "o":
					var (
						oe = (*OutputEvent)(e)
						// str = oe.String()
					)
					// consumed by reading <-AsciiTransportClient.OutputEvent()
					c.oech <- oe
					oe.Time = time.Since(c.start).Seconds()
					c.log(oe)
				default:
					log.Println("unknown message:", e)
				}
			}
			if line[0] == '{' {
				var (
					re  = new(ResizeEvent)
					err = json.Unmarshal(buf, re)
				)
				if err != nil {
					log.Println(err)
					continue
				}
				// var ( str = re.String() )
				// consumed by reading <-AsciiTransportServer.ResizeEvent()
				c.rech <- re
				re.Timestamp = uint(time.Now().Unix())
				c.log(re)
			}
		}
	}()
}

func (c *AsciiTransport) goWriteConn(w io.Writer) {
	var (
		clientInput2Server = func() {
			for {
				var (
					ie  = <-c.iech
					str = ie.String()
				)
				_, err := io.Copy(w, strings.NewReader(str))
				if err != nil {
					log.Println(err)
					break
				}
				ie.Time = time.Since(c.start).Seconds()
				c.log(ie)
			}
			c.Close()
		}
		clientResize2Server = func() {
			for {
				var (
					re  = <-c.rech
					str = re.String()
				)
				_, err := io.Copy(w, strings.NewReader(str))
				if err != nil {
					log.Println(err)
					break
				}
				re.Timestamp = uint(time.Now().Unix())
				c.log(re)
			}
			c.Close()
		}
		serverOutput2Client = func() {
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
	)
	if c.isClient {
		go clientInput2Server()
		go clientResize2Server()
		if c.reader != nil {
			go func() {
				for buf := make([]byte, 4096); ; {
					n, err := c.reader.Read(buf)
					if err != nil {
						log.Println(err)
						break
					}
					c.Input(buf[:n])
				}
				c.Close()
			}()
		}
		if c.writer != nil {
			go func() {
				for {
					oe := <-c.OutputEvent()
					_, err := io.Copy(c.writer, strings.NewReader(oe.Data))
					if err != nil {
						log.Println(err)
						break
					}
				}
				c.Close()
			}()
		}
	} else {
		go serverOutput2Client()
		if c.reader != nil {
			go func() {
				for buf := make([]byte, 4096); ; {
					n, err := c.reader.Read(buf)
					if err != nil {
						log.Println(err)
						break
					}
					c.Output(buf[:n])
				}
				c.Close()
			}()
		}
		if c.writer != nil {
			go func() {
				for {
					ie := <-c.InputEvent()
					_, err := io.Copy(c.writer, strings.NewReader(ie.Data))
					if err != nil {
						log.Println(err)
						break
					}
				}
				c.Close()
			}()
		}
	}
}
