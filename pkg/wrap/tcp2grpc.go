package wrap

import (
	"io"
	"net"
)

NewSingleListener(conn net.Conn, onAccept func()) net.Listener {
	return &singleListener{
		Conn: conn,
		OnAccept: onAccept,
	}
}

// single listener converts/upgrades the current tcp connection into grpc
// gender changer impl
type singleListener struct {
	net.Conn
	OnAccept func()
}

// singleListener implements the net.Listener interface
func (s *singleListener) Accept() (net.Conn, error) {
	if s.Conn != nil {
		s.OnAccept()
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
