package asciitransport

import "io"

type Opt func(at *AsciiTransport)

func WithReader(r io.Reader) Opt {
	return func(at *AsciiTransport) {
		at.src = r
	}
}

func WithWriter(w io.Writer) Opt {
	return func(at *AsciiTransport) {
		at.dst = w
	}
}

func WithResizeHook(rh func(_, _ uint16)) Opt {
	return func(at *AsciiTransport) {
		at.resizehook = rh
	}
}

func WithLogger(w io.WriteCloser) Opt {
	return func(at *AsciiTransport) {
		at.logger = NewLogger(w)
	}
}
