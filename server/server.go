package server

import (
	"log"
	"net/http"
	//"github.com/davecgh/go-spew/spew"
)

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
	}
}

func (s *Server) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ls)
	mux.HandleFunc("/ws", wslisten)
	mux.HandleFunc("/ws/", frontend)
	log.Println("listening on http://localhost" + s.Addr)
	return http.ListenAndServe(s.Addr, hijack(mux))
}

type Server struct {
	Addr string
}
