package main

import (
	"log"
	"os"

	"k0s.io/k0s/pkg/cli/wetty"
)

func main() {
	args := os.Args[1:]
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Fatalln(wetty.Run(args))
}
