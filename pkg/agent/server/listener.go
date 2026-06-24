package server

import (
	"net"
	"net/http"
	"sync"

	"github.com/btwiuse/wsconn"
)

// ChannelListener implements net.Listener
type ChannelListener struct {
	Conns    chan net.Conn
	closeOnce sync.Once
}

func (cl *ChannelListener) Accept() (net.Conn, error) {
	conn, ok := <-cl.Conns
	if !ok {
		return nil, net.ErrClosed
	}
	return conn, nil
}

func (cl *ChannelListener) Close() error {
	cl.closeOnce.Do(func() {
		close(cl.Conns)
	})
	return nil
}

func (cl *ChannelListener) Addr() net.Addr {
	return cl
}

func (cl *ChannelListener) Network() string {
	return "hijack"
}

func (cl *ChannelListener) String() string {
	return cl.Network()
}

func NewChannelListener() *ChannelListener {
	return &ChannelListener{
		Conns: make(chan net.Conn),
	}
}

type WSChannelListener struct {
	*ChannelListener
	Server *http.Server
	Conns  chan net.Conn
}

func NewWSChannelListener() *WSChannelListener {
	wl := &WSChannelListener{
		ChannelListener: NewChannelListener(),
		Server:          &http.Server{},
		Conns:           make(chan net.Conn),
	}
	wl.Server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := wsconn.Wrconn(w, r)
		if err != nil {
			return
		}
		wl.Conns <- conn
	})
	go wl.Server.Serve(wl.ChannelListener)
	return wl
}

func (wl *WSChannelListener) Accept() (net.Conn, error) {
	return <-wl.Conns, nil
}
