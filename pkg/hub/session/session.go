package session

import (
	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/conntroll/pkg/hub"
	"github.com/btwiuse/conntroll/pkg/uuid"
)

var (
	_ hub.Session = (*session)(nil)
)

type session struct {
	id string
	api.SessionClient
}

func (s *session) ID() string {
	return s.id
}

func NewSession(sc api.SessionClient) hub.Session {
	return &session{
		id:            uuid.New(),
		SessionClient: sc,
	}
}
