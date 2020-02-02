package main

import (
	"log"
	"os"

	"k0s.io/k0s/pkg/client/client"
	"k0s.io/k0s/pkg/client/config"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c := config.Parse(os.Args[1:])

	cl := client.NewClient(c)

	err := cl.Run()
	if err != nil {
		log.Println(err)
	}
}
