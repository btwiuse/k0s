package main

import (
	"os"

	"k0s.io/pkg/cli/chassis"
	"k0s.io/pkg/log"
)

func main() {
	chassis.Run(os.Args[1:])
}
