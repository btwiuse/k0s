package factory

import (
	"k0s.io/pkg/agent"
	"k0s.io/pkg/agent/tty"
)

func New(args []string) agent.TtyFactory {
	return &factory{args}
}

var (
	_ agent.TtyFactory = (*factory)(nil)
)

type factory struct {
	args []string
}

func (f *factory) MakeTty() (agent.Tty, error) {
	return tty.New(f.args)
}

func (f *factory) MakeTtyCmd(c []string) (agent.Tty, error) {
	return tty.New(c)
}

func (f *factory) MakeTtyEnv(c []string, env map[string]string) (agent.Tty, error) {
	return tty.NewEnv(c, env)
}
