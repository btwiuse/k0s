package config

// +build linux darwin
// +build !windows

import "os/exec"

func (c *config) getCmd() []string {
	shell := "bash"
	if _, err := exec.LookPath(shell); err != nil {
		shell = "sh"
	}
	args := []string{"env", "TERM=xterm", shell}
	if c.Cmd == "" {
		return args
	}
	return append(args, "-c", c.Cmd)
}
