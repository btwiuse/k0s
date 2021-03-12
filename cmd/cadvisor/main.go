package main

import (
	"log"
	"os"

	"k0s.io/pkg/cli/cadvisor"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := cadvisor.Run(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
}
