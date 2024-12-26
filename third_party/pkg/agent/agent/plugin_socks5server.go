//go:build plugin_socks5server

package server

import (
	"log"
	"net"

	"github.com/btwiuse/gost"
	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
)

func init() { Tunnels[api.Socks5] = StartSocks5Server }

func StartSocks5Server(c agent.Config) chan net.Conn {
	socks5Listener := NewChannelListener()
	go autoServe(socks5Listener)
	return socks5Listener.Conns
}

var autoHandler = gost.AutoHandler()

func autoServe(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println("socks5server: recovered from panic:", r, c)
				}
			}()
			autoHandler.Handle(c)
			c.Close()
		}()
	}
}
