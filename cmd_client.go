package main

import (
	"log"

	"github.com/btwiuse/conntroll/pkg/agent/agent"
	"github.com/btwiuse/conntroll/pkg/agent/config"
)

func clientCmd(args []string) {
	log.Println("TODO")
	return
	c := config.Parse(args)

	ag := agent.NewAgent(c)

	_ = ag
/*
	ag.Go(ag.ConnectAndServe)

	err := ag.Wait()
	if err != nil {
		log.Println()
	}
*/
}
