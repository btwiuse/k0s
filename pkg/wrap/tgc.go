// tcp gender changer impl. See https://en.wikipedia.org/wiki/TCP_Gender_Changer

package wrap

import (
	"io"
	"net"
)

func NewSingleListener(conn net.Conn, onAccept ...func()) net.Listener {
	return &singleListener{
		Conn:     conn,
		OnAccept: onAccept,
	}
}

// single listener help construct a grpc server from a tcp connection
type singleListener struct {
	net.Conn
	OnAccept []func()
}

// singleListener implements the net.Listener interface
func (s *singleListener) Accept() (net.Conn, error) {
	if s.Conn != nil {
		if len(s.OnAccept) != 0 {
			s.OnAccept[0]()
		}
		c := s.Conn
		s.Conn = nil
		return c, nil
	}
	return nil, io.EOF
}

func (s *singleListener) Close() error {
	return nil
}

func (s *singleListener) Addr() net.Addr {
	return s.Conn.LocalAddr()
}
