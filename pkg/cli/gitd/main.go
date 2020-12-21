package gitd

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"k0s.io/k0s/pkg/tunnel/listener"
	"k0s.io/k0s/pkg/middleware"
)

var (
    serviceHandler = dumpHandler
    uploadPackHandler = dumpHandler
    receivePackHandler = dumpHandler
)

func Run(args []string) (err error) {
    var (
        port = args[0]
        ln = listener.Listener(port, "/")
        gitMux = newGitMux()
    )

	log.Println("server is listening on", port)

	return http.Serve(ln, middleware.LoggingMiddleware(gitMux))
}

func dumpHandler(w http.ResponseWriter, r *http.Request){
    log.Println(r.Method, r.RequestURI)
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