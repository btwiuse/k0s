package dialer

import (
	"github.com/btwiuse/conntroll/pkg/agent"
)

var (
	_ agent.Dialer = (*dialr)(nil)
)

func New(c agent.Config) agent.Dialer {
	return &dialr{
		c: c,
	}
}

type dialr struct {
	c agent.Config
}
