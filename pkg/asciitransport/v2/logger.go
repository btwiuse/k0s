package asciitransport

import (
	"io"
	"log"
)

func WithLogger(w io.WriteCloser) Opt {
	return func(at *AsciiTransport) {
		at.logger = NewLogger(w)
	}
}

type Logger interface {
	Print(v interface{})
	Close() error
}

func NewLogger(w io.WriteCloser) Logger {
	l := &logger{
		l: log.New(w, "", 0),
		w: w,
	}
	return l
}

type logger struct {
	l *log.Logger
	w io.WriteCloser
}

func (l *logger) Print(v interface{}) {
	l.l.Print(v)
}

func (l *logger) Close() error {
	return l.w.Close()
}
