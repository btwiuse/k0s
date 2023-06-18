package hub

import (
	"log"
	"net/http"

	"k0s.io/pkg/hub/config"
	"k0s.io/pkg/hub/server"
	"k0s.io/pkg/tunnel/listener"
)

func Run2(args []string) (err error) {
	c := config.Parse(args)

	ln := listener.Listener(c.Port(), "/")

	log.Println("server is listening on", c.Port())

	h := server.NewHub(c)

	if c.UseTLS() {
		err = http.ServeTLS(ln, h.Handler(), c.Cert(), c.Key())
	} else {
		err = http.Serve(ln, h.Handler())
	}

	return
}
