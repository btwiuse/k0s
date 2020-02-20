package main

import (
	"log"
	"os"

	"k0s.io/k0s/pkg/cli/hub"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	hub.Run(os.Args[1:])
}
