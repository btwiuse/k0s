package wrap

import (
	"io"
)

type ReadWriteCloser struct {
	Reader io.Reader
	Writer io.Writer
	Closer io.Closer
}

func WrapReadWriteCloser(r io.Reader, wc io.WriteCloser) io.ReadWriteCloser {
	return &ReadWriteCloser{
		Reader: r,
		Writer: wc,
		Closer: wc,
	}
}

func (b *ReadWriteCloser) Close() error {
	return b.Closer.Close()
}

func (b *ReadWriteCloser) Write(p []byte) (int, error) {
	return b.Writer.Write(p)
}

func (b *ReadWriteCloser) Read(p []byte) (int, error) {
	return b.Reader.Read(p)
}
