package agent

import (
	"net"

	"github.com/ginuerzh/gost"
	types "k0s.io/pkg/agent"
)

func StartRedirectServer(c types.Config) chan net.Conn {
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
