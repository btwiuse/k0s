package main

import (
	"log"
	"time"

	"github.com/btwiuse/conntroll/pkg/agent/agent"
	"github.com/btwiuse/conntroll/pkg/agent/config"
)

func agentCmd(args []string) {
	c := config.Parse(args)

	ag := agent.NewAgent(c)

	for range time.Tick(time.Second) {
		err := ag.ConnectAndServe()
		if err != nil {
			log.Println(err, "Reconnect")
		}
	}
}
