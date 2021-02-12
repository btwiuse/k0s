package hub

import (
	"log"
	"net"
	"net/http"

	"k0s.io/k0s/pkg/wrap"
)

var (
	_ http.Handler = (*lys)(nil)
	_ net.Listener = (*lys)(nil)
)

// lys is an http handler to net.Listener converter
// lys implements net.Listener/http.Handler
type lys struct {
	conns chan net.Conn
}

func (l *lys) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.Header)
	conn, err := wrap.Wrconn(w, r)
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
		addr := wrap.NewAddr(network, hostport)
		conn = wrap.ConnWithAddr(conn, addr)
	}

	l.conns <- conn
}

func (l *lys) Accept() (net.Conn, error) {
	return <-l.conns, nil
}

func (l *lys) Close() error {
	return nil
}

func (l *lys) Addr() net.Addr {
	return l
}

func (l *lys) Network() string {
	return "hijack"
}

func (l *lys) String() string {
	return l.Network()
}

func NewHandleHijackListener() *lys {
	return &lys{conns: make(chan net.Conn)}
}
