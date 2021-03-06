package config

func (c *Config) getCmd() []string {
	shell := "cmd"
	args := []string{shell}
	if c.Cmd == "" {
		return args
	}
	return append(args, "/c", c.Cmd)
}
