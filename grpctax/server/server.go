package main

import (
	"errors"
	"log"
	"net"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/btwiuse/invctrl/grpctax/tax"
)

func init() {
	grpclog.SetLogger(log.New(os.Stderr, "", 0))
}

type taxServer struct {
	rate float64
}

func (s *taxServer) Tax(ctx context.Context, amt *tax.Amount) (*tax.Amount, error) {
	if amt.Cents < 0 {
		return nil, errors.New("Negative amount not allowed")
	}
	return &tax.Amount{Cents: int32(float64(amt.Cents)*s.rate + .5)}, nil
}

func main() {
	grpclog.Println("listening on tcp://127.0.0.1:1234")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		grpclog.Fatalf(err.Error())
	}
	grpcServer := grpc.NewServer()
	tax.RegisterTaxComputerServer(grpcServer, &taxServer{.05})
	grpclog.Fatalln(grpcServer.Serve(listener))
}
