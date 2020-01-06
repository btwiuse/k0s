package session

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/conntroll/pkg/hub"
	"github.com/btwiuse/conntroll/pkg/uuid"
	"google.golang.org/grpc"
)

var (
	_ hub.Session = (*session)(nil)
)

type session struct {
	id   string
	name string
	api.SessionClient
	created time.Time
	quit    chan struct{}
	onclose func()
}

func (s *session) Time() time.Time {
	return s.created
}

func (s *session) Name() string {
	return s.name
}

func (s *session) ID() string {
	return s.id
}

// TODO
func (s *session) Close() error {
	close(s.quit)
	return nil
}

func (s *session) gc() {
	select {
	case <-s.quit:
		s.onclose()
	}
}

func NewSession(name string, conn net.Conn) hub.Session {
	var (
		cc *grpc.ClientConn  = toGRPCClientConn(conn)
		sc api.SessionClient = api.NewSessionClient(cc)
		s                    = &session{
			id:            uuid.New(),
			name:          name,
			SessionClient: sc,
			created:       time.Now(),
			quit:          make(chan struct{}),
			onclose:       func() { cc.Close() },
		}
	)
	go s.gc()
	return s
}

func toGRPCClientConn(c net.Conn) *grpc.ClientConn {
	options := []grpc.DialOption{
		// grpc.WithTransportCredentials(creds),
		grpc.WithInsecure(),
		grpc.WithContextDialer(
			func(ctx context.Context, s string) (net.Conn, error) {
				return c, nil
			},
		),
	}

	// TODO: handle this
	cc, err := grpc.Dial("", options...)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cc
}
