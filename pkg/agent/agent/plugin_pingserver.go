package agent

import (
	"net"

	types "k0s.io/pkg/agent"
	"k0s.io/pkg/api"
)

func init() { Tunnels[api.Ping] = StartPingServer }

func StartPingServer(c types.Config) chan net.Conn {
	pingListener := NewLys()
	go pingServe(pingListener)
	return pingListener.Conns
}

func pingServe(l net.Listener) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		go func() {
			defer c.Close()
			// nop
		}()
	}
}
