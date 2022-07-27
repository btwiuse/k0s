//go:build plugin_jsonlserver

package server

import (
	"bufio"
	"log"
	"net"

	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
)

func init() {
	Tunnels[api.Jsonl] = StartJsonlServer
	println("JSONL")
}

func StartJsonlServer(c agent.Config) chan net.Conn {
	var (
		ro            bool     = c.GetReadOnly()
		defaultCmd    []string = c.GetCmd()
		jsonlListener          = NewLys()
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
