package main

import "github.com/btwiuse/jiri/cmdline"

var cmdRoot = &cmdline.Command{
	Name:     "conntroll",
	LookPath: true,
	Children: []*cmdline.Command{
		cmdHub,
		cmdAgent,
		cmdClient,
	},
}

func main() {
	cmdline.Main(cmdRoot)
}
