package wrap

import (
	"io"
	"net"
	"net/http"
)

func Hijack(w http.ResponseWriter) (net.Conn, error) {
	return HijackConn(w.(http.Hijacker).Hijack())
}

func HijackConn(conn net.Conn, buf io.Reader, err error) (net.Conn, error) {
	return &Conn{
		Conn: conn,
		Buf:  buf,
	}, err
}

type Conn struct {
	net.Conn
	Buf io.Reader
}

func (c *Conn) Read(b []byte) (int, error) {
	return io.MultiReader(c.Buf, c.Conn).Read(b)
}
