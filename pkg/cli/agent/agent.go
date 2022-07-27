package agent

import (
	"fmt"
	"log"

	"k0s.io/pkg/agent/config"
	"k0s.io/pkg/agent/server"
)

func Run(args []string) error {
	c := config.Parse(args)

	log.Println(fmt.Sprintf("agent name: %s (%s)", c.GetName(), c.GetID()))

	ag := server.NewAgent(c)

	err := ag.ConnectAndServe()

	return err
}
