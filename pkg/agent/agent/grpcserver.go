package agent

import (
	"net/http"

	"google.golang.org/grpc"
	types "k0s.io/conntroll/pkg/agent"
	"k0s.io/conntroll/pkg/agent/tty"
	"k0s.io/conntroll/pkg/api"
	"k0s.io/conntroll/pkg/api/grpcimpl"
	"k0s.io/conntroll/pkg/exporter"
)

func StartGrpcServer(c types.Config) types.GrpcServer {
	var (
		cmd []string = c.GetCmd()
		ro  bool     = c.GetReadOnly()
	)
	grpcListener := NewLys()
	grpcServer := grpc.NewServer()
	api.RegisterSessionServer(grpcServer, &grpcimpl.Session{
		ReadOnly:       ro,
		TtyFactory:     tty.NewFactory(cmd),
		FileServer:     http.FileServer(http.Dir("/")),
		MetricsHandler: exporter.NewHandler(),
	})
	go grpcServer.Serve(grpcListener)
	return grpcListener
}
