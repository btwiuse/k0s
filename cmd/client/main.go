package main

import (
	"os"

	"k0s.io/pkg/cli/client"
	"k0s.io/pkg/log"
)

func main() {
	if err := client.Run(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}
