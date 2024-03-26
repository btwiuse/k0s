package tunnel

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/btwiuse/wsconn"
)

func NewTunnel() *Tunnel {
	return &Tunnel{
		cc:   make(chan net.Conn),
		quit: make(chan struct{}),
		once: &sync.Once{},
	}
}

type Tunnel struct {
	cc   chan net.Conn
	quit chan struct{}
	once *sync.Once
}

var _ net.Listener = (*Tunnel)(nil)

func (t *Tunnel) Accept() (net.Conn, error) {
	select {
	case conn := <-t.cc:
		return conn, nil
	case <-t.quit:
		return nil, fmt.Errorf("tunnel closed")
	}
}

func (t *Tunnel) Close() error {
	t.once.Do(func() {
		close(t.quit)
	})
	return nil
}

var _ net.Addr = (*Tunnel)(nil)

func (t *Tunnel) Addr() net.Addr {
	return t
}

func (t *Tunnel) Network() string {
	return "tunnel"
}

func (t *Tunnel) String() string {
	return t.Network()
}

var _ http.Handler = (*Tunnel)(nil)

func (t *Tunnel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := wsconn.Wrconn(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	t.Push(conn)
}

func (t *Tunnel) Push(conn net.Conn) {
	t.cc <- conn
}
