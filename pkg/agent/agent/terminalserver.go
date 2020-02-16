package agent

import (
	"log"
	"net"

	types "k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/agent/tty"
	"k0s.io/k0s/pkg/asciitransport"
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
				asciitransport.WithWriter(term),
			}
			server := asciitransport.Server(conn, opts...)
			// send
			// case output:

			// recv
			go func() {
				for {
					var (
						re   = <-server.ResizeEvent()
						rows = int(re.Height)
						cols = int(re.Width)
					)
					err := term.Resize(rows, cols)
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
