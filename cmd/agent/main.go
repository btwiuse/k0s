package main

import (
	"log"
	"os"
	"time"

	"github.com/btwiuse/conntroll/pkg/agent/agent"
	"github.com/btwiuse/conntroll/pkg/agent/config"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c := config.Parse(os.Args[1:])

	ag := agent.NewAgent(c)

	for range time.Tick(time.Second) {
		ag.Go(ag.ConnectAndServe)

		err := ag.Wait()
		if err != nil {
			log.Println()
		}
	}
}
