package impl

import (
	"log"
	"net/url"

	"github.com/btwiuse/wetty/pkg/localcmd"
	"github.com/btwiuse/wetty/pkg/utils"
	"google.golang.org/grpc"

	"github.com/btwiuse/conntroll/pkg/agent/config"
	"github.com/btwiuse/conntroll/pkg/agent/dial"
	"github.com/btwiuse/conntroll/pkg/api"
	grpcimpl "github.com/btwiuse/conntroll/pkg/api/grpc/impl"
)

type NewSession struct {
	*localcmd.Factory
	Name string
}

type NewSessionRequest struct {
	Info url.Values
}

type NewSessionResponse struct{}

func (c *NewSession) New(req NewSessionRequest, res *NewSessionResponse) error {
	log.Println("NewSession.New called with", req)

	conn, err := dial.WithInfo(config.Default.Server, req.Info)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	api.RegisterSessionServer(grpcServer, &grpcimpl.Session{Factory: c.Factory, Name: c.Name})
	grpcServer.Serve(&utils.SingleListener{conn})
	return nil
}
