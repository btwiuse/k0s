package main

import (
	"log"
	"net/rpc"

	"github.com/google/uuid"
	"github.com/btwiuse/conntroll/pkg/agent/config"
	"github.com/btwiuse/conntroll/pkg/agent/dial"
	rpcimpl "github.com/btwiuse/conntroll/pkg/api/rpc/impl"
	"github.com/btwiuse/wetty/localcmd"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	config.Init()

	conn, err := dial.WithInfo(config.Default.Server, config.Default.Info)
	if err != nil {
		log.Fatalln("err connecting to hub:", err)
	}

	log.Println(config.Default)
	log.Println("connected as:", config.Default.Info.Get("id"))

	factory := &localcmd.Factory{
		Args: []string{"/usr/bin/env", "TERM=xterm", "bash"},
	}

	rpcServer := rpc.NewServer()
	rpcServer.Register(&rpcimpl.NewSlave{Factory: factory, Name: uuid.New().String()})
	rpcServer.ServeConn(conn)
}
