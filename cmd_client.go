package main

import (
	"github.com/btwiuse/jiri"
	"github.com/btwiuse/jiri/cmdline"
)

var cmdClient = &cmdline.Command{
	Runner: jiri.RunnerFunc(runClient),
	Name:   "client",
}

func runClient(jirix *jiri.X, args []string) error {
	return nil
}
