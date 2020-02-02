package main

import (
	"log"

	"k0s.io/k0s/pkg/hub/config"
	"k0s.io/k0s/pkg/hub/hub"
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
