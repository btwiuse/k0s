package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/btwiuse/invctrl/protocol"
)

func main() {
	req := protocol.Request{
		Command: "hello",
	}
	resp := new(protocol.Response)

	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}
	rpcClient := rpc.NewClient(conn)

	/* this works the same
	rpcClient, err := rpc.Dial("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}
	*/

	if err := rpcClient.Call("Hello.Execute", req, resp); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Message)
}
