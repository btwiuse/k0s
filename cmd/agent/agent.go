package main

import (
	"log"
	"net/rpc"

	"github.com/btwiuse/invctrl/pkg/agent/config"
	"github.com/btwiuse/invctrl/pkg/agent/dial"
	rpcimpl "github.com/btwiuse/invctrl/pkg/api/rpc/impl"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	c := config.Init()
	conn := dial.Dial(c)
	c.Id = dial.Handshake(conn)
	config.Default = c

	/* subsequent call
	        conn := dialRemote(config.Default)
		handshake(conn)
	*/

	log.Println("connected:", c.Id)

	rpcServer := rpc.NewServer()
	rpcServer.Register(new(rpcimpl.Conn))
	rpcServer.Register(new(rpcimpl.WsConn))
	rpcServer.Register(new(rpcimpl.GRPCConn))
	log.Println("serveconn")
	rpcServer.ServeConn(conn)
	log.Fatalln("bye")
}
