package ui

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/webteleport/utils"
)

func withUI(ui http.Handler, r *mux.Router) *mux.Router {
	r.NotFoundHandler = ui
	return r
}

// NewRouter create a new gorilla *mux.Router, with NotFoundHandler set to reverse proxy of `url`
func NewRouter(url string) *mux.Router {
	return withUI(utils.ReverseProxy(url), mux.NewRouter())
}
