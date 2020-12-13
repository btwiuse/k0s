package main

import (
	"log"
	"os"

	"k0s.io/k0s/pkg/cli/hub"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Fatalln(hub.Run(os.Args[1:]))
}
