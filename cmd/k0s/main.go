package main

import (
	"os"

	"k0s.io/pkg/cli/k0s"
	"k0s.io/pkg/log"
)

func main() {
	log.Fatalln(k0s.Run(os.Args[1:]))
}
