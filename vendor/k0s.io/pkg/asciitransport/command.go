package asciitransport

func WithCommand(c []string) Opt {
	return func(at *AsciiTransport) {
		at.cmd = c
	}
}

func WithEnv(env map[string]string) Opt {
	return func(at *AsciiTransport) {
		at.env = env
	}
}
