package main

import (
	"log"
	"os"

	"k0s.io/pkg/cli/pcp"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := pcp.Run(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
}
