package main

import (
	"log"
	"os"

	"k0s.io/pkg/cli/wetty"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Fatalln(wetty.Run(os.Args[1:]))
}
