// +build linux darwin
// +build !windows

package config

import "os/exec"

func (c *config) getCmd() []string {
	shell := "bash"
	if _, err := exec.LookPath(shell); err != nil {
		shell = "sh"
	}
	args := []string{"/usr/bin/env", "TERM=xterm", shell}
	if c.Cmd == "" {
		return args
	}
	return append(args, "-c", c.Cmd)
}
