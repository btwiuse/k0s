package asciitransport

import "io"

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
