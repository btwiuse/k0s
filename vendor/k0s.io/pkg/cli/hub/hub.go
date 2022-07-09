package hub

import (
	"fmt"
	"log"

	"k0s.io/pkg/hub/config"
	"k0s.io/pkg/hub/hub"
	"k0s.io/pkg/hub/self"
)

func Run(args []string) (err error) {
	c := config.Parse(args)

	log.Println("hub is listening on", c.Port())

	h := hub.NewHub(c)

	go self.Agent(fmt.Sprintf("http://127.0.0.1%s", c.Port()))

	if c.UseTLS() {
		err = h.ListenAndServeTLS(c.Cert(), c.Key())
	} else {
		err = h.ListenAndServe()
	}

	return
}
