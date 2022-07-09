package chassis

import (
	"log"
	"net/http"

	"k0s.io/pkg/simple/listener"
	"k0s.io/pkg/tunnel/handler"
)

func Run(args []string) (err error) {
	if len(args) < 1 {
		log.Fatalln("wrong number of args")
	}
	port := args[0]
	log.Println("Listening on", port)
	log.Println(http.Serve(listener.Listener(port), handler.Handler("/")))
	return nil
}
