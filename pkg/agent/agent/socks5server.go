package agent

import (
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/riobard/go-shadowsocks2/socks"
	types "k0s.io/k0s/pkg/agent"
)

func StartSocks5Server(c types.Config) types.Socks5Server {
	socks5Listener := NewLys()
	go serve(socks5Listener)
	return socks5Listener
}

func serve(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			logf("failed to accept: %v", err)
			continue
		}

		go func() {
			defer c.Close()
			tcpKeepAlive(c)
			// c = shadow(c)

			// tgt, err := socks.ReadAddr(c)
			tgt, err := socks.Handshake(c)
			if err != nil {
				logf("failed to get target address: %v", err)
				return
			}

			rc, err := net.Dial("tcp", tgt.String())
			if err != nil {
				logf("failed to connect to target: %v", err)
				return
			}
			defer rc.Close()
			tcpKeepAlive(rc)

			logf("proxy %s <-> %s", c.RemoteAddr(), tgt)
			if err = relay(c, rc); err != nil {
				if err, ok := err.(net.Error); ok && err.Timeout() {
					return // ignore i/o timeout
				}
				logf("relay error: %v", err)
			}
		}()
	}
}

func tcpKeepAlive(c net.Conn) {
	if tcp, ok := c.(*net.TCPConn); ok {
		tcp.SetKeepAlive(true)
		tcp.SetKeepAlivePeriod(3 * time.Minute)
	}
}

func logf(f string, v ...interface{}) {
	// if config.Verbose {
	log.Printf(f, v...)
	// }
}

// relay copies between left and right bidirectionally. Returns any error occurred.
func relay(left, right net.Conn) error {
	var err, err1 error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err1 = io.Copy(right, left)
		right.SetReadDeadline(time.Now()) // unblock read on right
	}()

	_, err = io.Copy(left, right)
	left.SetReadDeadline(time.Now()) // unblock read on left
	wg.Wait()

	if err1 != nil {
		err = err1
	}
	return err
}
