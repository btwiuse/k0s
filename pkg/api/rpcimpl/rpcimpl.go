// +build rpcimpl

package rpcimpl

import (
	"log"

	"github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/conntroll/pkg/api/grpcimpl"
	"github.com/btwiuse/conntroll/pkg/wrap"
	"google.golang.org/grpc"
)

type RPC struct {
	Agent agent.Agent
}

type RPCRequest struct{}

type RPCResponse struct{}

func (c *RPC) New(req RPCRequest, res *RPCResponse) error {
	c.Agent.Go(c.Agent.ConnectAndServe)
	log.Println("RPC.New invoked")
	return nil
}

type Session struct {
	Agent agent.Agent
}

type SessionRequest struct{}

type SessionResponse struct{}

func (c *Session) New(req SessionRequest, res *SessionResponse) error {
	conn, err := c.Agent.CreateSession()
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	api.RegisterSessionServer(grpcServer, &grpcimpl.Session{
		TtyFactory: c.Agent,
	})
	onAccept := func() {
		log.Println("Gender Change: TCP Client -> gRPC Server")
	}
	grpcServer.Serve(wrap.NewSingleListener(conn, onAccept))
	return nil
}
