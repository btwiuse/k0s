package main

import (
	"k0s.io/conntroll/pkg/client/client"
	"k0s.io/conntroll/pkg/client/config"
)

func clientCmd(args []string) {
	c := config.Parse(args)

	cl := client.NewClient(c)

	cl.Run()
}
