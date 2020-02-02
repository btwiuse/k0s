package terminal

import (
	"net"
	"time"

	"github.com/btwiuse/asciitransport"
	"k0s.io/k0s/pkg/hub"
	"k0s.io/k0s/pkg/uuid"
)

var (
	_ hub.Terminal = (*terminal)(nil)
)

type terminal struct {
	id      string
	name    string
	conn    net.Conn
	created time.Time
}

func (s *terminal) Time() time.Time {
	return s.created
}

func (s *terminal) Name() string {
	return s.name
}

func (s *terminal) ID() string {
	return s.id
}

// TODO
func (s *terminal) Close() error {
	return s.AsciiTransportClient.Close()
}

func NewTerminal(conn net.Conn) net.Conn {
	var (
		opts = []asciitransport.Opt{
			// asciitransport.WithReader()
		}
		s = &terminal{
			id:                   uuid.New(),
			name:                 name,
			AsciiTransportClient: asciitransport.Client(conn),
			created:              time.Now(),
		}
	)
	go s.gc()
	return s
}
