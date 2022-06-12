// +build windows

package tty

import (
        "github.com/ActiveState/termtest/conpty"
        "k0s.io/pkg/agent"
        "syscall"
        "os"
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
        *conpty.ConPty
}

func (t *term) Read(buf []byte) (int, error) {
        return t.ConPty.OutPipe().Read(buf)
}

func (t *term) Write(buf []byte) (int, error) {
        u32, err := t.ConPty.Write(buf)
        return int(u32), err
}

func (t *term) Resize(rows, cols int) error {
        return t.ConPty.Resize(uint16(cols), uint16(rows))
}

type factory struct {
        args []string
}

func (f *factory) MakeTty() (agent.Tty, error) {
        pty, err := conpty.New(80, 24)
        if err != nil {
                return nil, err
        }
        pid, _, err := pty.Spawn(f.args[0], []string{}, &syscall.ProcAttr{Env: os.Environ()})
        if err != nil {
                return nil, err
        }
        _ = pid
        return &term{ConPty: pty}, nil
}
