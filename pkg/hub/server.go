package hub

import (
	"net/http"
)

func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws/", static)
	mux.HandleFunc("/agents/", getAgents)
	mux.HandleFunc("/new/agent", newAgent)
	mux.HandleFunc("/new/slave", newSlave)
	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
