//go:build linux || darwin || freebsd || openbsd || netbsd
// +build linux darwin freebsd openbsd netbsd

package tty

import (
	"io"
	"os"
	"os/exec"
	"syscall"
	"unsafe"

	"github.com/creack/pty"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"k0s.io/pkg/agent"
)

var _ agent.Tty = (*term)(nil)

type term struct {
	args      []string
	process   *os.Process
	pty       *os.File
	ptyReader io.Reader
	ptyWriter io.Writer
	ptyClosed chan struct{}
}

func (t *term) Close() error {
	if t.process != nil {
		t.process.Signal(syscall.SIGINT)
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
	sz := &pty.Winsize{
		uint16(rows),
		uint16(cols),
		0,
		0,
	}
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		t.pty.Fd(),
		syscall.TIOCSWINSZ,
		uintptr(unsafe.Pointer(sz)),
	)
	if errno != 0 {
		return errno
	}
	return nil
}

func New(args []string) (agent.Tty, error) {
	return NewEnv(args, map[string]string{})
}

func NewEnv(args []string, env map[string]string) (agent.Tty, error) {
	envArray := append([]string{"TERM=xterm"}, map2arr(env)...)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = append(os.Environ(), envArray...)

	pty, err := pty.Start(cmd)
	if err != nil {
		// todo close cmd?
		return nil, err // ors.Wrapf(err, "failed to start command `%s`", command)
	}
	ptyClosed := make(chan struct{})

	t := &term{
		args:      args,
		process:   cmd.Process,
		pty:       pty,
		ptyReader: transform.NewReader(pty, unicode.UTF8.NewDecoder()),
		ptyWriter: transform.NewWriter(pty, unicode.UTF8.NewEncoder()),
		ptyClosed: ptyClosed,
	}

	// When the process is closed by the user,
	// close pty so that Read() on the pty breaks with an EOF.
	go func() {
		defer func() {
			t.pty.Close()
			close(t.ptyClosed)
		}()

		t.process.Wait()
	}()

	t.Resize(24, 80)
	return t, nil
}
