package hub

import (
	"log"
	"net"
	"net/http"
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
	log.Println(r.Header)
	conn, err := wrconn(w, r)
	if err != nil {
		log.Println("error ws accept:", err)
		return
	}

	if forward := r.Header.Get("X-Forwarded-For"); forward != "" {
		addr := NewAddr(conn.RemoteAddr().Network(), forward)
		conn = ConnWithAddr(conn, addr)
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
