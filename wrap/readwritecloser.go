package wrap

import (
	"io"
)

type ReadWriteCloser struct {
	Reader      io.Reader
	WriteCloser io.WriteCloser
}

func WrapReadWriteCloser(r io.Reader, wc io.WriteCloser) *ReadWriteCloser {
	return &ReadWriteCloser{
		Reader:      r,
		WriteCloser: wc,
	}
}

func (b *ReadWriteCloser) Close() error {
	return b.WriteCloser.Close()
}

func (b *ReadWriteCloser) Write(p []byte) (int, error) {
	return b.WriteCloser.Write(p)
}

func (b *ReadWriteCloser) Read(p []byte) (int, error) {
	return b.Reader.Read(p)
}
