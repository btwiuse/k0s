package client

import (
	types "github.com/btwiuse/conntroll/pkg/client"
)

type client struct {
}

func NewClient(c client.Config) types.Client {
	return &client{}
}

func (c *client) ListAgents() []types.Agent {
	// curl /api/agents/
	// return &agent{}
	return nil
}
