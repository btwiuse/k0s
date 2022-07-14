package main

import (
	"os"

	"k0s.io/pkg/cli/hub"
	"k0s.io/pkg/log"
)

func main() {
	if err := hub.Run2(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}
