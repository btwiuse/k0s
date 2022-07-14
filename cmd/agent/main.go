package main

import (
	"os"

	"k0s.io/pkg/cli/agent"
	"k0s.io/pkg/log"
)

func main() {
	if err := agent.Run(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}
