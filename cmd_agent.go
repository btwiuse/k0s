package main

import (
	"log"
	"time"

	"k0s.io/conntroll/pkg/agent/agent"
	"k0s.io/conntroll/pkg/agent/config"
)

func agentCmd(args []string) {
	c := config.Parse(args)

	ag := agent.NewAgent(c)

	for range time.Tick(time.Second) {
		err := ag.ConnectAndServe()
		if err != nil {
			log.Println(err)
		}
	}
}
