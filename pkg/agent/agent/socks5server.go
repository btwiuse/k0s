package agent

import (
	socks5 "github.com/armon/go-socks5"
	types "k0s.io/conntroll/pkg/agent"
)

func StartSocks5Server(c types.Config) types.Socks5Server {
	socks5Listener := NewLys()
	emptyConfig := &socks5.Config{}
	socks5Server, _ := socks5.New(emptyConfig)
	go socks5Server.Serve(socks5Listener)
	return socks5Listener
}
