package session

import (
	"time"

	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/conntroll/pkg/hub"
	"github.com/btwiuse/conntroll/pkg/uuid"
)

var (
	_ hub.Session = (*session)(nil)
)

type session struct {
	id   string
	name string
	api.SessionClient
	created time.Time
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

func NewSession(name string, sc api.SessionClient) hub.Session {
	return &session{
		id:            uuid.New(),
		name:          name,
		SessionClient: sc,
		created:       time.Now(),
	}
}
