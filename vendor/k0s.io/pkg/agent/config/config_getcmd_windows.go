package config

func (c *Config) getCmd() []string {
	shell := `C:\WINDOWS\System32\WindowsPowerShell\v1.0\powershell.exe`
	args := []string{shell}
	if c.Cmd == "" {
		return args
	}
	return append(args, "-Command", c.Cmd)
}
