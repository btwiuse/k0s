package main

import (
	"log"
	"os"

	"k0s.io/pkg/cli/kubectl"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := kubectl.Run(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
}
