// +build yrpc

package yrpcimpl

import (
	"github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/conntroll/pkg/api/grpcimpl"
	"github.com/yrpc/yrpc"
	"google.golang.org/grpc"
)

const (
	NewSessionRequest yrpc.Cmd = iota
	NewSessionResponse
)

type YServer struct {
	Agent agent.Agent
}

// traditional tgc: instantiate a new grpcServer for each connection
// new tgc: create connection and send to net.Listener
// yrpc mux handle (SessionRequest, func(w, r))
func (ys *YServer) NewSession() error {
	conn, err := ys.Agent.CreateSession() // send /api/agent/grpc?id=xxxx-xxxx
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	api.RegisterSessionServer(grpcServer, &grpcimpl.Session{
		TtyFactory: ys.Agent,
	})

	/*
		onAccept := func() {
			log.Println("Gender Change: TCP Client -> gRPC Server")
		}
		grpcServer.Serve(wrap.NewSingleListener(conn, onAccept))
	*/

	return nil
}
