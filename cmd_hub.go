package main

import (
	"log"

	"github.com/btwiuse/conntroll/pkg/hub/config"
	"github.com/btwiuse/conntroll/pkg/hub/hub"
)

func hubCmd(args []string) {
	c := config.Parse(args)

	log.Println("server is listening on", c.Port())

	h := hub.NewHub(c)

	if c.UseTLS() {
		log.Fatalln(h.ListenAndServeTLS(c.Cert(), c.Key()))
	}

	log.Fatalln(h.ListenAndServe())
}
