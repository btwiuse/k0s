package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"k0s.io/conntroll/pkg/agent/agent"
	"k0s.io/conntroll/pkg/agent/config"
)

func init() {
	go func() {
		if http.ListenAndServe(":31337", nil) != nil {
			os.Exit(0)
		}
	}()
}

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
