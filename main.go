package main

import (
	"os"

	"k0s.io/pkg/cli/k0s"
	"k0s.io/pkg/log"
)

func main() {
	if err := k0s.Run(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}
