package main

import (
	"log"
	"os"

	"k0s.io/pkg/cli/gos"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Fatalln(gos.Run(os.Args[1:]))
}
