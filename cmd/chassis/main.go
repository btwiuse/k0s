package main

import (
	"log"
	"net/http"
	"os"

	"k0s.io/k0s/pkg/simple/listener"
	"k0s.io/k0s/pkg/tunnel/handler"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("wrong number of args")
	}
	port := os.Args[1]
	log.Println("Listening on", port)
	log.Println(http.Serve(listener.Listener(port), handler.Handler("/")))
}
