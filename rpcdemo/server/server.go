package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/btwiuse/invctrl/protocol"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("server listening on :8000")

	// rpc.Register(new(protocol.Hello))
	rpcServer := rpc.NewServer()
	rpcServer.Register(new(protocol.Hello))

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("new conn:", conn.RemoteAddr())
		go rpcServer.ServeConn(conn)
	}
}
