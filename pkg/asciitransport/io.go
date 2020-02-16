package asciitransport

import "io"

type Resizer interface {
	Resize(height, width uint16)
}

func WithReader(r io.Reader) Opt {
	return func(at *AsciiTransport) {
		at.reader = r
	}
}

func WithWriter(w io.Writer) Opt {
	return func(at *AsciiTransport) {
		at.writer = w
	}
}

func WithResizer(Resizer) Opt {
	return func(at *AsciiTransport) {
	}
}
