package main

import (
	"log"
	"net/rpc"

	"github.com/btwiuse/wetty/localcmd"
	"github.com/btwiuse/invctrl/pkg/agent/config"
	"github.com/btwiuse/invctrl/pkg/agent/dial"
	rpcimpl "github.com/btwiuse/invctrl/pkg/api/rpc/impl"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	c := config.Init()
	conn := dial.DialAgent(c.Server)

	log.Println("connected:", c.Id)

	factory := &localcmd.Factory{
		Args: []string{"/usr/bin/env", "TERM=xterm", "bash"},
	}

	rpcServer := rpc.NewServer()
	rpcServer.Register(&rpcimpl.NewSlave{factory})
	log.Println("serveconn")
	rpcServer.ServeConn(conn)
	log.Fatalln("bye")
}
