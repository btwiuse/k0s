package main

import (
	"log"
	"net/http"
	"strings"

	ethrpc "github.com/ethereum/go-ethereum/rpc"
	// "github.com/gdr3941/websocketrpc"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterService(new(HelloService), "")
	// http.Handle("/", s)
}

func handleAll() http.Handler {
	hello := new(HelloService)
	erpc := ethrpc.NewServer()
	erpc.RegisterName("hello", hello)

	wsHandler := erpc.WebsocketHandler([]string{"*"})
	httpHandler := erpc

	handler := func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		isWS := strings.Split(r.Header.Get("Upgrade"), ",")[0] == "websocket"
		switch {
		case isWS:
			wsHandler.ServeHTTP(w, r)
		default:
			httpHandler.ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(handler)
}

func main() {
	log.Println(":8000")
	//log.Fatalln(http.ListenAndServe(":8000", nil))
	// rpcs := websocketrpc.NewServer(":8000", true)
	// rpcs.Register(new(HelloService))
	// rpcs.StartServer()

	log.Fatalln(http.ListenAndServe(":8000", handleAll()))
}

type HelloArgs struct {
	Who string
}

type HelloReply struct {
	Message string
}

type HelloService struct{}

func (h *HelloService) Say2(args *HelloArgs, reply *HelloReply) error {
	reply.Message = "Hello, " + args.Who + "!"
	return nil
}

func (h *HelloService) Say(args HelloArgs) (reply HelloReply) {
	reply.Message = "Hello, " + args.Who + "!"
	return
}
