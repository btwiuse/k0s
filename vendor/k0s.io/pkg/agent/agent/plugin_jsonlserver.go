//go:build plugin_jsonlserver
// +build plugin_jsonlserver

package agent

import (
	"bufio"
	"log"
	"net"

	types "k0s.io/pkg/agent"
	"k0s.io/pkg/api"
)

func init() {
	Tunnels[api.Jsonl] = StartJsonlServer
	println("JSONL")
}

func StartJsonlServer(c types.Config) chan net.Conn {
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
