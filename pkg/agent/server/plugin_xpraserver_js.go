//go:build js

package server

import (
	"net"

	"k0s.io/pkg/agent/config"
)

func StartXpraServer(c *config.Config) chan net.Conn {
	ch := make(chan net.Conn)
	go func() {
		for conn := range ch {
			conn.Close()
		}
	}()
	return ch
}
