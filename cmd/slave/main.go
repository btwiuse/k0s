package main

import (
	"log"
	"net/rpc"

	"github.com/btwiuse/invctrl/pkg/slave/config"
	"github.com/btwiuse/invctrl/pkg/slave/dial"
	"github.com/btwiuse/invctrl/protocol"
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
	rpcServer.Register(new(protocol.Hello))
	rpcServer.Register(new(protocol.Bash))
	rpcServer.Register(new(protocol.Conn))
	rpcServer.Register(new(protocol.WsConn))
	rpcServer.Register(new(protocol.Rootfs))
	rpcServer.Register(new(protocol.Echo))
	rpcServer.Register(new(protocol.GRPCConn))
	log.Println("serveconn")
	rpcServer.ServeConn(conn)
	log.Fatalln("bye")
}
