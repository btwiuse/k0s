package hub

import (
	"log"

	"k0s.io/pkg/hub/config"
	"k0s.io/pkg/hub/hub"
)

func Run(args []string) (err error) {
	c := config.Parse(args)

	log.Println("server is listening on", c.Port())

	h := hub.NewHub(c)

	if c.UseTLS() {
		err = h.ListenAndServeTLS(c.Cert(), c.Key())
	} else {
		err = h.ListenAndServe()
	}

	return
}
