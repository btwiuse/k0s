//go:build raw && !gorilla && !nhooyr
// +build raw,!gorilla,!nhooyr

package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
)

func (cl *client) dial(p string, h http.Header) (conn net.Conn, err error) {
	var (
		c  = cl.Config
		ub = &url.URL{
			Scheme: c.GetSchemeWS(),
			Host:   c.GetAddr(),
			Path:   p,
		}
		reqURI = ub.RequestURI()
	)

	switch c.GetScheme() {
	case "http":
		conn, err = net.Dial("tcp", c.GetAddr())
		if err != nil {
			return nil, err
		}
	case "https":
		conn, err = tls.Dial("tcp", c.GetAddr(), &tls.Config{
			InsecureSkipVerify: c.GetInsecure(),
		})
	}

	if err != nil {
		return nil, err
	}

	_, err = io.WriteString(conn, fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nConnection: Keep-Alive\r\n\r\n", reqURI, c.GetHost()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
