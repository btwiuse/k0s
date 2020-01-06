package agent

import (
	// "net/rpc"
	// "time"

	"github.com/btwiuse/conntroll/pkg"
	"github.com/btwiuse/conntroll/pkg/hub"
	"github.com/btwiuse/conntroll/pkg/manager"
)

var (
	_ hub.SessionManager = (*sessionManager)(nil)
)

type sessionManager struct {
	pkg.Manager
}

func (sm *sessionManager) AddSession(s hub.Session) {
	sm.Manager.Add(s)
}

func (sm *sessionManager) GetSession(id string) hub.Session {
	return sm.Get(id).(hub.Session)
}

func NewSessionManager() hub.SessionManager {
	return &sessionManager{
		Manager: manager.NewManager(),
	}
}
