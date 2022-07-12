package main

import (
	"os"

	"k0s.io/pkg/cli/client"
	"k0s.io/pkg/log"
)

func main() {
	log.Fatalln(client.Run(os.Args[1:]))
}
