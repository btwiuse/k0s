package hub

import (
	"fmt"
	"log"
	"net/http"

	"github.com/webteleport/webteleport/ufo"

	"k0s.io/pkg/hub/config"
	"k0s.io/pkg/hub/self"
	"k0s.io/pkg/hub/server"
)

func Run(args []string) (err error) {
	c := config.Parse(args)

	log.Println("hub is listening on", c.Port())

	h := server.NewHub(c)

	go self.Agent(fmt.Sprintf("http://127.0.0.1%s", c.Port()))

	if c.Ufo() != "" {
		go ufo.Serve(c.Ufo(), h.Handler())
	}

	if c.UseTLS() {
		err = http.ListenAndServeTLS(c.Port(), c.Cert(), c.Key(), h.Handler())
	} else {
		err = http.ListenAndServe(c.Port(), h.Handler())
	}

	return
}
