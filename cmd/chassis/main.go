package main

import (
	"os"

	"k0s.io/pkg/cli/chassis"
)

func main() {
	chassis.Run(os.Args[1:])
}
