package server

import (
	"log"
	"net/http"
	//"github.com/davecgh/go-spew/spew"
)

func NewServer(addr string) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ls)
	mux.HandleFunc("/ws", wslisten)
	mux.HandleFunc("/ws/", frontend)
	return &Server{
		Addr:    addr,
		Handler: hijack(mux),
	}
}

func (s *Server) Run() error {
	log.Println("listening on http://localhost" + s.Addr)
	return http.ListenAndServe(s.Addr, s.Handler)
}

type Server struct {
	Addr    string
	Handler http.Handler
}
