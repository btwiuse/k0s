package main

import (
	"log"
	"os"

	"k0s.io/third_party/pkg/cli/wetty"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Fatalln(wetty.Run(os.Args[1:]))
}
