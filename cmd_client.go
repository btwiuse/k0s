package main

import (
	"log"

	"k0s.io/conntroll/pkg/client/client"
	"k0s.io/conntroll/pkg/client/config"
)

func clientCmd(args []string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c := config.Parse(args)

	cl := client.NewClient(c)

	err := cl.Run()
	if err != nil {
		log.Println(err)
	}
}
