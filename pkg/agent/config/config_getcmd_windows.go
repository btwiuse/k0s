package config

// +build windows
// +build !linux
// +build !darwin

func (c *config) getCmd() []string {
	shell := "cmd"
	args := []string{shell}
	if c.Cmd == "" {
		return args
	}
	return append(args, "/c", c.Cmd)
}
