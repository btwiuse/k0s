package server

import (
	"bufio"
	"log"
	"net"

	"k0s.io/pkg/agent/config"
)

func StartJsonlServer(c *config.Config) chan net.Conn {
	var (
		ro            bool     = c.ReadOnly
		defaultCmd    []string = c.GetCmd()
		jsonlListener          = NewWSChannelListener()
	)
	_ = ro
	go serveJsonl(jsonlListener, defaultCmd)
	return jsonlListener.Conns
}

func serveJsonl(ln net.Listener, defaultCmd []string) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			scanner := bufio.NewScanner(conn)
			defer conn.Close()
			for scanner.Scan() {
				log.Println(scanner.Text())
			}
		}()
	}
}
