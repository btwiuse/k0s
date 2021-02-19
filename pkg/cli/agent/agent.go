package agent

import (
	"fmt"
	"log"

	"k0s.io/pkg/agent/agent"
	"k0s.io/pkg/agent/config"
)

func Run(args []string) error {
	c := config.Parse(args)

	log.Println(fmt.Sprintf("agent name: %s (%s)", c.GetName(), c.GetID()))

	ag := agent.NewAgent(c)

	err := ag.ConnectAndServe()

	return err
}
