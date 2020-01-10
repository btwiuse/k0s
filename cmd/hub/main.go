package main

import (
	"log"
	"os"

	//"github.com/davecgh/go-spew/spew"

	"k0s.io/conntroll/pkg/hub/config"
	"k0s.io/conntroll/pkg/hub/hub"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c := config.Parse(os.Args[1:])

	log.Println("server is listening on", c.Port())

	h := hub.NewHub(c)

	if c.UseTLS() {
		log.Fatalln(h.ListenAndServeTLS(c.Cert(), c.Key()))
	}

	log.Fatalln(h.ListenAndServe())
}
