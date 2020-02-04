package main

import (
	"log"

	"k0s.io/k0s/pkg/client/client"
	"k0s.io/k0s/pkg/client/config"
)

func clientCmd(args []string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c := config.Parse(args)

	cl := client.NewClient(c)

	for {
		err := cl.Run()
		if err != nil {
			log.Println(err)
		}
	}
}
