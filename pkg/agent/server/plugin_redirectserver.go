//go:build plugin_redirectserver

package server

import (
	"net"

	"github.com/btwiuse/gost"
	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
)

func init() { Tunnels[api.Redir] = StartRedirectServer }

func StartRedirectServer(c agent.Config) chan net.Conn {
	redirectListener := NewLys()
	go redirectServe(redirectListener)
	return redirectListener.Conns
}

var redirectHandler = gost.TCPRedirectHandler()

func redirectServe(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		go func() {
			defer c.Close()
			redirectHandler.Handle(c)
		}()
	}
}
