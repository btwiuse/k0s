package main

import (
	"fmt"

	"github.com/btwiuse/jiri"
	"github.com/btwiuse/jiri/cmdline"
)

var cmdHub = &cmdline.Command{
	Runner: jiri.RunnerFunc(runHub),
	Name:   "hub",
}

var hubFlags struct {
	portFlag string
}

func init() {
	flags := &cmdHub.Flags
	flags.StringVar(&hubFlags.portFlag, "port", ":8000", "hub listening port")
}

func runHub(jirix *jiri.X, args []string) error {
	fmt.Println(args)
	return nil
}
