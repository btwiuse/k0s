package server

import (
	"k0s.io"
	"k0s.io/pkg/hub"
	"k0s.io/pkg/manager"
)

var (
	_ hub.AgentManager = (*agentManager)(nil)
)

type agentManager struct {
	k0s.Manager
}

func (am *agentManager) AddAgent(ag hub.Agent) {
	am.Manager.Add(ag)
}

func (am *agentManager) GetAgents() []hub.Agent {
	var (
		all = am.Manager.Values()
		ret = []hub.Agent{}
	)
	for _, v := range all {
		ret = append(ret, v.(hub.Agent))
	}
	return ret
}

func (am *agentManager) GetAgent(id string) hub.Agent {
	return am.Manager.Get(id).(hub.Agent)
}

func NewAgentManager() hub.AgentManager {
	return &agentManager{
		Manager: manager.NewManager(),
	}
}
