// +build windows
// +build !linux

package tty

import (
	"github.com/liamg/aminal/platform"
	"k0s.io/conntroll/pkg/agent"
)

func NewFactory(args []string) agent.TtyFactory {
	return &factory{
		args: args,
	}
}

var (
	_ agent.TtyFactory = (*factory)(nil)
	_ agent.Tty        = (*term)(nil)
)

type term struct {
	platform.Pty
}

func (t *term) Resize(rows, cols int) error {
	return t.Pty.Resize(cols, rows)
}

type factory struct {
	args []string
}

func (f *factory) MakeTty() (agent.Tty, error) {
	pty, err := platform.NewPty(80, 24)
	if err != nil {
		return nil, err
	}
	proc, err := pty.CreateGuestProcess(f.args[0])
	if err != nil {
		return nil, err
	}
	_ = proc
	return &term{Pty: pty}, nil
}
