//go:build js

package server

import (
	"net"

	"k0s.io/pkg/agent/config"
)

func StartXpraServer(c *config.Config) chan net.Conn {
	return make(chan net.Conn)
}
