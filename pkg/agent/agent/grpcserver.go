package agent

import (
	"google.golang.org/grpc"
	types "k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/agent/tty"
	"k0s.io/k0s/pkg/api"
	"k0s.io/k0s/pkg/api/grpcimpl"
)

func StartGrpcServer(c types.Config) types.GrpcServer {
	var (
		cmd []string = c.GetCmd()
		ro  bool     = c.GetReadOnly()
	)
	grpcListener := NewLys()
	grpcServer := grpc.NewServer()
	api.RegisterSessionServer(grpcServer, &grpcimpl.Session{
		ReadOnly:   ro,
		TtyFactory: tty.NewFactory(cmd),
	})
	go grpcServer.Serve(grpcListener)
	return grpcListener
}
