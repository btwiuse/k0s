package main

import (
	"log"

	"github.com/btwiuse/conntroll/pkg/agent/agent"
	"github.com/btwiuse/conntroll/pkg/agent/config"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c := config.Parse()

	ag := agent.NewAgent(c)

	ag.Go(ag.ConnectAndServe)

	log.Println(ag.Wait())
}
