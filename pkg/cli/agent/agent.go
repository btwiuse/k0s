package agent

import (
    "fmt"
    "log"

	"k0s.io/k0s/pkg/agent/agent"
	"k0s.io/k0s/pkg/agent/config"
)

func Run(args []string) error {
	c := config.Parse(args)

	log.Println(fmt.Sprintf("agent name: %s (%s)", c.GetName(), c.GetID()))

	ag := agent.NewAgent(c)

	err := ag.ConnectAndServe()

    return err
}
