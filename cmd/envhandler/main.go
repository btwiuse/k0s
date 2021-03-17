package main

import (
	"log"
	"net"
	"net/http"
	"k0s.io/pkg/exporter/env"
)

func main(){
	port := ":8000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	handler := env.NewHandler()
	log.Println("listening on", port)
	log.Fatalln(http.Serve(listener, handler))
}
