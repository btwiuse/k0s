// +build !gorilla !nhooyr

package dialer

import (
	"fmt"
	"io"
	"net"
)

func (d *dialr) Dial(p string) (conn net.Conn, err error) {
	var (
		c = d.c
	)

	conn, err = net.Dial("tcp", c.GetAddr())
	if err != nil {
		return nil, err
	}
	_, err = io.WriteString(conn, fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nConnection: Keep-Alive\r\n\r\n", p, c.GetHostname()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
