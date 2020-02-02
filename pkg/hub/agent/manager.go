package agent

import (
	// "net/rpc"
	// "time"

	"k0s.io/k0s/pkg"
	"k0s.io/k0s/pkg/hub"
	"k0s.io/k0s/pkg/manager"
)

var (
	_ hub.SessionManager  = (*sessionManager)(nil)
	_ hub.TerminalManager = (*terminalManager)(nil)
)

type terminalManager struct {
	pkg.Manager
}

func (sm *terminalManager) AddTerminal(s hub.Terminal) {
	sm.Manager.Add(s)
}

func (sm *terminalManager) GetTerminal(id string) hub.Terminal {
	return sm.Get(id).(hub.Terminal)
}

func NewTerminalManager() hub.TerminalManager {
	return &terminalManager{
		Manager: manager.NewManager(),
	}
}

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
