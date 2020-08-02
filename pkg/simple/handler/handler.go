package handler

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"
)

func Handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", rootfs)
	return handlers.LoggingHandler(os.Stderr, cors.AllowAll().Handler(r))
}

func rootfs(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("/usr/local/go")).ServeHTTP(w, r)
}
