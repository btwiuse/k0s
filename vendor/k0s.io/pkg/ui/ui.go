package ui

import (
	"net/http"

	"github.com/gorilla/mux"
	"k0s.io/pkg/reverseproxy"
)

func reverseHandler(url string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reverseproxy.Handler(url).ServeHTTP(w, r)
	})
}

func withUI(ui http.Handler, r *mux.Router) *mux.Router {
	r.NotFoundHandler = ui
	return r
}

// NewRouter create a new gorilla *mux.Router, with NotFoundHandler set to reverse proxy of `url`
func NewRouter(url string) *mux.Router {
	return withUI(reverseHandler(url), mux.NewRouter())
}
