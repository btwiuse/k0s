package main

import (
	"os"

	"k0s.io/pkg/cli/agent"
	"k0s.io/pkg/log"
)

func main() {
	log.Fatalln(agent.Run(os.Args[1:]))
}
