package main

import (
	"log"
	"os"
	"time"

	"k0s.io/k0s/pkg/agent/agent"
	"k0s.io/k0s/pkg/agent/config"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c := config.Parse(os.Args[1:])

	ag := agent.NewAgent(c)

	for range time.Tick(time.Second) {
		err := ag.ConnectAndServe()
		if err != nil {
			log.Println(err)
		}
	}
}
