package hub

import (
	// "io"
	"net"
	"net/rpc"

	"github.com/btwiuse/conntroll/pkg/api"
)

type IDer interface {
	ID() string
	// CreationTime() time.Time
}

type Manager interface {
	Add(IDer)
	Del(string)
	Has(string) bool
	Get(string) IDer
	Values() []IDer
	Size() int
}

type Hub interface {
	// IDer
	AgentManager

	// Serve(net.Listener) error
	ListenAndServe() error
}

type AgentManager interface {
	Manager

	AddAgent(Agent)
	GetAgent(string) Agent
	GetAgents() []Agent
}

type Agent interface {
	IDer
	SessionManager

	AddRPCConn(net.Conn)
	AddSessionConn(net.Conn)

	NewRPC()
	NewSession() Session
}

type SessionManager interface {
	Manager

	AddSession(Session)
	GetSession(string) Session
}

type Session interface {
	IDer
	api.SessionClient

	// TTY() io.ReaderFrom // | io.WriterTo
	// FS(api.ChunkRequest) io.ReaderFrom
}

type RPC interface {
	IDer

	// *rpc.Client
	Call(serviceMethod string, args interface{}, reply interface{}) error
	Close() error
	Go(serviceMethod string, args interface{}, reply interface{}, done chan *rpc.Call) *rpc.Call
}

type RPCManager interface {
	Manager

	AddRPC(RPC)
	Last() RPC
}
