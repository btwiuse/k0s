package asciitransport

import (
	"io"
	"log"

	"google.golang.org/protobuf/encoding/protojson"

	"k0s.io/third_party/pkg/asciiproto"
)

type Logger interface {
	Log(*asciiproto.Frame)
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

func (l *logger) Log(f *asciiproto.Frame) {
	b, _ := protojson.Marshal(f)
	l.l.Println(string(b))
}

func (l *logger) Close() error {
	return l.w.Close()
}
