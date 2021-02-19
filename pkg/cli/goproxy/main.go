package goproxy

import (
	"log"
	"net/http"

	"github.com/goproxy/goproxy"

	"k0s.io/pkg/simple/listener"
	"k0s.io/pkg/middleware"
)

func Run(args []string) (err error) {
	port := ":8000"
	if len(args) > 1 {
		port = args[0]
	}
	log.Println("Listening on", port)
	return http.Serve(
		listener.Listener(port),
		middleware.LoggingMiddleware(goproxy.New()),
	)
}
