package main

import (
	"os"

	"k0s.io/pkg/cli/hub"
	"k0s.io/pkg/log"
)

func main() {
	log.Fatalln(hub.Run(os.Args[1:]))
}
