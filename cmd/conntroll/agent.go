package main

import (
	"github.com/btwiuse/jiri"
	"github.com/btwiuse/jiri/cmdline"
)

var cmdAgent = &cmdline.Command{
	Runner: jiri.RunnerFunc(runAgent),
	Name:   "agent",
}

func runAgent(jirix *jiri.X, args []string) error {
	return nil
}
