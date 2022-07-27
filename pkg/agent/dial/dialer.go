package dial

import (
	"k0s.io/pkg/agent"
)

var (
	_ agent.Dialer = (*dialer)(nil)
)

func New(c agent.Config) agent.Dialer {
	return &dialer{
		c: c,
	}
}

type dialer struct {
	c agent.Config
}
