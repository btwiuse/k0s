package agent

import (
	"log"
	"net"
	"os"

	"github.com/btwiuse/asciitransport"
	types "k0s.io/conntroll/pkg/agent"
	"k0s.io/conntroll/pkg/agent/tty"
)

func StartTerminalServer(c types.Config) types.TerminalServer {
	var (
		cmd              []string = c.GetCmd()
		ro               bool     = c.GetReadOnly()
		fac                       = tty.NewFactory(cmd)
		terminalListener          = NewLys()
	)
	_ = ro
	go serveTerminal(terminalListener, fac)
	return terminalListener
}

func serveTerminal(ln net.Listener, fac types.TtyFactory) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			term, err := fac.MakeTty()
			if err != nil {
				log.Println(err)
				return
			}

			opts := []asciitransport.Opt{
				asciitransport.WithReader(term),
				asciitransport.WithReader(term),
				asciitransport.WithLogger(os.Stderr),
			}
			server := asciitransport.Server(conn, opts...)
			// send
			// case output:

			// recv
			go func() {
				for {
					var (
						re     = <-server.ResizeEvent()
						width  = int(re.Width)
						height = int(re.Height)
					)
					err := term.Resize(width, height)
					if err != nil {
						log.Println(err)
						break
					}
				}
				server.Close()
			}()
		}()
	}
}

/*
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
*/
