package hub

import (
	"net/http"

	"github.com/rs/cors"
)

func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/agent/", static)
	mux.HandleFunc("/api/agents/", getAgents)
	mux.HandleFunc("/api/new", newAgentOrSession)
	return &http.Server{
		Addr:    addr,
		Handler: cors.Default().Handler(mux),
	}
}
