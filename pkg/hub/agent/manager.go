package agent

import (
	"net/rpc"

	"github.com/btwiuse/conntroll/pkg/hub"
	"github.com/btwiuse/conntroll/pkg/manager"
	"github.com/btwiuse/conntroll/pkg/uuid"
)

var (
	_ hub.SessionManager = (*sessionManager)(nil)
	_ hub.RPCManager     = (*rpcManager)(nil)
)

type sessionManager struct {
	hub.Manager
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

type rpcManager struct {
	hub.Manager
}

// todo: use iterator
func (rm *rpcManager) Last() hub.RPC {
	return rm.Values()[rm.Size()-1].(hub.RPC)
}

// todo: use iterator
func (rm *rpcManager) AddRPC(rc hub.RPC) {
	if rm.Size() > 3 {
		// rm the first
	}
	rm.Manager.Add(rc)
}

func NewRPCManager() hub.RPCManager {
	return &rpcManager{
		Manager: manager.NewManager(),
	}
}

type rpcc struct {
	id string
	*rpc.Client
}

func (r *rpcc) ID() string {
	return r.id
}

func ToRPC(rc *rpc.Client) hub.RPC {
	return &rpcc{
		id:     uuid.New(),
		Client: rc,
	}
}
