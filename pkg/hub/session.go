package hub

import (
	"net/rpc"

	"github.com/btwiuse/conntroll/pkg/api"
	rpcimpl "github.com/btwiuse/conntroll/pkg/api/rpc/impl"
)

type Session struct {
	api.SessionClient
}

func NewSession(agent *Agent) *Session {
	session := &Session{}
	req := rpcimpl.NewSessionRequest{}
	resp := new(rpcimpl.NewSessionResponse)

	done := make(chan *rpc.Call, 1)
	agent.RPCClient.Go("NewSession.New", req, resp, done)
	<-done
	cc := <-agent.GRPCClientConn
	session.SessionClient = api.NewSessionClient(cc)
	return session
}
