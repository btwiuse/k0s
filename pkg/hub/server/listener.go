package server

import (
	"log"
	"net"
	"net/http"

	"github.com/btwiuse/wsconn"
	"k0s.io/pkg/agent/server"
)

var (
	_ http.Handler = (*HTTPChannelListener)(nil)
	_ net.Listener = (*HTTPChannelListener)(nil)
)

// HTTPChannelListener is an http handler to net.Listener converter
// HTTPChannelListener implements net.Listener/http.Handler
type HTTPChannelListener struct {
	*server.ChannelListener
}

func (l *HTTPChannelListener) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println("HTTPChannelListener.ServeHTTP")
	// log.Println(r.Header)
	conn, err := wsconn.Wrconn(w, r)
	if err != nil {
		log.Println("error ws accept:", err)
		return
	}

	if forwardIP := r.Header.Get("X-Forwarded-For"); forwardIP != "" {
		forwardPort := r.Header.Get("X-Forwarded-Port")
		if forwardPort == "" {
			forwardPort = "0"
		}
		network := conn.RemoteAddr().Network()
		hostport := forwardIP + ":" + forwardPort
		addr := wsconn.NewAddr(network, hostport)
		conn = wsconn.ConnWithAddr(conn, addr)
	}

	l.ChannelListener.Conns <- conn
}

func NewHTTPChannelListener() *HTTPChannelListener {
	return &HTTPChannelListener{
		server.NewChannelListener(),
	}
}
