package impl

import (
	"io"
	"log"
	"net"
	"net/url"

	"github.com/btwiuse/wetty/localcmd"
	"google.golang.org/grpc"

	"github.com/btwiuse/invctrl/pkg/agent/config"
	"github.com/btwiuse/invctrl/pkg/agent/dial"
	"github.com/btwiuse/invctrl/pkg/api"
	grpcimpl "github.com/btwiuse/invctrl/pkg/api/grpc/impl"
)

type NewSlave struct{
	*localcmd.Factory
}

type NewSlaveRequest struct {
	Args url.Values
}

type NewSlaveResponse struct {
	Message string
}

func (c *NewSlave) New(req NewSlaveRequest, res *NewSlaveResponse) error {
	log.Println("NewSlave.New called with", req)

	res.Message = "OK"
	conn := dial.DialSlave(config.Default.HttpServer, req.Args)
	go func() {
		grpcServer := grpc.NewServer()
		api.RegisterBidiStreamServer(grpcServer, &grpcimpl.BidiStream{c.Factory})
		grpcServer.Serve(&singleListener{conn})
		log.Println("grpcServer exited")
	}()
	return nil
}

type singleListener struct {
	conn net.Conn
}

// singleListener implements the net.Listener interface
func (s *singleListener) Accept() (net.Conn, error) {
	log.Println("Accept")
	if c := s.conn; c != nil {
		s.conn = nil
		return c, nil
	}
	return nil, io.EOF
}

func (s *singleListener) Close() error {
	return nil
}

func (s *singleListener) Addr() net.Addr {
	return s.conn.LocalAddr()
}
