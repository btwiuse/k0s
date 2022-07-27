package hub

import (
	"log"

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
		err = h.ServeTLS(ln, c.Cert(), c.Key())
	} else {
		err = h.Serve(ln)
	}

	return
}
