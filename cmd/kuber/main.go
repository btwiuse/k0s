package main

import (
	"os"

	"k0s.io/pkg/cli/kuber"
	"k0s.io/pkg/log"
)

func main() {
	if err := kuber.Run(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}
