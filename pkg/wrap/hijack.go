package wrap

import (
	"io"
	"net"
	"net/http"
	"time"
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
	Conn net.Conn
	Buf  io.Reader
}

func (c *Conn) Read(b []byte) (int, error) {
	return io.MultiReader(c.Buf, c.Conn).Read(b)
}

func (c *Conn) Write(b []byte) (int, error) {
	return c.Conn.Write(b)
}

func (c *Conn) Close() error {
	return c.Conn.Close()
}

func (c *Conn) LocalAddr() net.Addr {
	return c.Conn.LocalAddr()
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Conn) SetDeadline(t time.Time) error {
	return c.Conn.SetDeadline(t)
}

func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.Conn.SetReadDeadline(t)
}

func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.Conn.SetWriteDeadline(t)
}

/*
func Hijack() (net.Conn, *bufio.ReadWriter, error)

type Conn interface {
	Read(b []byte) (n int, err error)
	Write(b []byte) (n int, err error)
	Close() error
	LocalAddr() Addr
	RemoteAddr() Addr
	SetDeadline(t time.Time) error
	SetReadDeadline(t time.Time) error
	SetWriteDeadline(t time.Time) error
}
*/
