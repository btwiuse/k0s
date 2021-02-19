package main

import (
	"log"
	"os"

	"k0s.io/pkg/cli/buildkite"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Fatalln(buildkite.Run(os.Args[1:]))
}
