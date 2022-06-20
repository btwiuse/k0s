//go:build windows
// +build windows

package tty

import (
	"io"
	"os"
	"syscall"

	"github.com/ActiveState/termtest/conpty"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"k0s.io/pkg/agent"
)

var _ agent.Tty = (*term)(nil)

type term struct {
	args []string
	cpty *conpty.ConPty

	process *os.Process
	// pty       *os.File
	ptyReader io.Reader
	ptyWriter io.Writer
	ptyClosed chan struct{}
}

func (t *term) Close() error {
	if t.process != nil {
		t.process.Signal(syscall.SIGINT)
		t.process.Kill()
	}
	for {
		select {
		case <-t.ptyClosed:
			return nil
		}
	}
}

func (t *term) Write(p []byte) (int, error) {
	return t.ptyWriter.Write(p)
}

func (t *term) Read(p []byte) (int, error) {
	return t.ptyReader.Read(p)
}

func (t *term) Resize(rows, cols int) error {
	return t.cpty.Resize(uint16(cols), uint16(rows))
}

func New(args []string) (agent.Tty, error) {
	return NewEnv(args, map[string]string{})
}

func NewEnv(args []string, env map[string]string) (agent.Tty, error) {
	cpty, err := conpty.New(80, 24)
	if err != nil {
		return nil, err
	}

	envArray := append([]string{"TERM=xterm"}, map2arr(env)...)
	attr := &syscall.ProcAttr{
		Env: append(os.Environ(), envArray...),
	}

	pid, _, err := cpty.Spawn(args[0], args[1:], attr)
	if err != nil {
		return nil, err
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}

	ptyClosed := make(chan struct{})

	t := &term{
		args:      args,
		process:   process,
		cpty:      cpty,
		ptyReader: transform.NewReader(cpty.OutPipe(), unicode.UTF8.NewDecoder()),
		ptyWriter: transform.NewWriter(cpty.InPipe(), unicode.UTF8.NewEncoder()),
		ptyClosed: ptyClosed,
	}

	// When the process is closed by the user,
	// close pty so that Read() on the pty breaks with an EOF.
	go func() {
		defer func() {
			t.cpty.Close()
			close(t.ptyClosed)
		}()

		t.process.Wait()
	}()

	t.Resize(24, 80)
	return t, nil
}
