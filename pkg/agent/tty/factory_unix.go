// +build linux darwin freebsd openbsd netbsd

package tty

import (
	"k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/localcmd"
)

func NewFactory(args []string) agent.TtyFactory {
	return &factory{
		f: &localcmd.Factory{
			Args: args,
		},
	}
}

var (
	_ agent.TtyFactory = (*factory)(nil)
	_ agent.Tty        = (*term)(nil)
)

type term struct {
	Lc *localcmd.Lc
}

func (t *term) Close() error {
	return t.Lc.Close()
}

func (t *term) Write(p []byte) (int, error) {
	return t.Lc.Write(p)
}

func (t *term) Read(p []byte) (int, error) {
	return t.Lc.Read(p)
}

func (t *term) Resize(rows, cols int) error {
	return t.Lc.Resize(rows, cols)
}

type factory struct {
	f *localcmd.Factory
}

func (f *factory) MakeTty() (agent.Tty, error) {
	// f.f.Args = args
	lc, err := f.f.New()
	if err != nil {
		return nil, err
	}
	return &term{
		Lc: lc,
	}, nil
}
