package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/btwiuse/invctrl/grpctax/tax"
)

func init() {
	grpclog.SetLogger(log.New(os.Stderr, "", 0))
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf(err.Error())
	}
	defer conn.Close()
	client := tax.NewTaxComputerClient(conn)
	amt := &tax.Amount{Cents: 300}
	tax, err := client.Tax(context.Background(), amt)
	if err != nil {
		grpclog.Fatalf(err.Error())
	}
	fmt.Println("Tax on", amt.Cents, "cents is", tax.Cents, "cents")
}
