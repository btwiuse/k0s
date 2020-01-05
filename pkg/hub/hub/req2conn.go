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
	conn, err := wrconn(w, r)
	if err != nil {
		log.Println("error ws accept:", err)
		return
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
