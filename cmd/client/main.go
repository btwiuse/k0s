package main

import (
	"log"
	"os"

	"k0s.io/conntroll/pkg/client/client"
	"k0s.io/conntroll/pkg/client/config"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c := config.Parse(os.Args[1:])

	cl := client.NewClient(c)

	cl.Run()
}
