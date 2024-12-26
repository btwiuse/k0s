package server

import (
	"net"
	"net/http"

	"github.com/btwiuse/wsconn"
)

// ChannelListener implements net.Listener
type ChannelListener struct {
	Conns chan net.Conn
}

func (cl *ChannelListener) Accept() (net.Conn, error) {
	return <-cl.Conns, nil
}

func (cl *ChannelListener) Close() error {
	// TODO: close Conns
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
