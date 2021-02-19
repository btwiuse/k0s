package gitd

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"k0s.io/pkg/middleware"
)

func Serve(ln net.Listener) error {
	var (
		gitMux  = newGitMux()
		handler = middleware.LoggingMiddleware(gitMux)
	)

	return http.Serve(ln, handler)
}

func newGitMux() http.Handler {
	r := mux.NewRouter()

	// shamelessly borrowed from
	// https://github.com/pratikju/servidor/blob/master/server.go#L28

	r.HandleFunc(`/info/refs`, serviceHandler).Methods("GET")
	r.HandleFunc(`/git-upload-pack`, uploadPackHandler).Methods("POST")
	r.HandleFunc(`/git-receive-pack`, receivePackHandler).Methods("POST")

	return r
}
