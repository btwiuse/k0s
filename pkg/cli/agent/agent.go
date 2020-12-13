package agent

import (
	"k0s.io/k0s/pkg/agent/agent"
	"k0s.io/k0s/pkg/agent/config"
)

func Run(args []string) (err error) {
	c := config.Parse(args)

	ag := agent.NewAgent(c)

	return ag.ConnectAndServe()
}
