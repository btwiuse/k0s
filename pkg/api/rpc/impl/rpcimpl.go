package impl

import (
	"log"
	"net/url"

	"github.com/btwiuse/wetty/localcmd"
	"github.com/btwiuse/wetty/utils"
	"google.golang.org/grpc"

	"github.com/btwiuse/invctrl/pkg/agent/config"
	"github.com/btwiuse/invctrl/pkg/agent/dial"
	"github.com/btwiuse/invctrl/pkg/api"
	grpcimpl "github.com/btwiuse/invctrl/pkg/api/grpc/impl"
)

type NewSlave struct {
	*localcmd.Factory
	Name string
}

type NewSlaveRequest struct {
	Info url.Values
}

type NewSlaveResponse struct{}

func (c *NewSlave) New(req NewSlaveRequest, res *NewSlaveResponse) error {
	log.Println("NewSlave.New called with", req)

	conn, err := dial.WithInfo(config.Default.Server, req.Info)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	api.RegisterSlaveServer(grpcServer, &grpcimpl.Slave{Factory: c.Factory, Name: c.Name})
	grpcServer.Serve(&utils.SingleListener{conn})
	return nil
}
