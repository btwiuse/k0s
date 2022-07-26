package hub

import (
	"k0s.io"
	types "k0s.io/pkg/hub"
	"k0s.io/pkg/manager"
)

var (
	_ types.AgentManager = (*agentManager)(nil)
)

type agentManager struct {
	k0s.Manager
}

func (am *agentManager) AddAgent(ag types.Agent) {
	am.Manager.Add(ag)
}

func (am *agentManager) GetAgents() []types.Agent {
	var (
		all = am.Manager.Values()
		ret = []types.Agent{}
	)
	for _, v := range all {
		ret = append(ret, v.(types.Agent))
	}
	return ret
}

func (am *agentManager) GetAgent(id string) types.Agent {
	return am.Manager.Get(id).(types.Agent)
}

func NewAgentManager() types.AgentManager {
	return &agentManager{
		Manager: manager.NewManager(),
	}
}
