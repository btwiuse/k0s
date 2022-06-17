package agent

import (
	"log"
	"net"
	"sync"

	types "k0s.io/pkg/agent"
	"k0s.io/pkg/agent/tty/factory"
	"k0s.io/pkg/api"
	"k0s.io/pkg/asciitransport"
)

func init() { Tunnels[api.Terminal] = StartTerminalServer }

func StartTerminalServer(c types.Config) chan net.Conn {
	var (
		ro               bool     = c.GetReadOnly()
		defaultCmd       []string = c.GetCmd()
		terminalListener          = NewLys()
	)
	_ = ro
	go serveTerminal(terminalListener, defaultCmd)
	return terminalListener.Conns
}

func serveTerminal(ln net.Listener, defaultCmd []string) {
	var fac types.TtyFactory = factory.New(defaultCmd)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			var (
				tryCommandOnce = &sync.Once{}
				cmdCh          = make(chan []string, 1)
				resizeCh       = make(chan struct{ rows, cols int }, 4)
			)

			server := asciitransport.Server(conn)
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
					tryCommandOnce.Do(func() {
						cmdCh <- re.Command
					})
					resizeCh <- struct{ rows, cols int }{rows, cols}
				}
				server.Close()
			}()

			cmd := <-cmdCh

			if len(cmd) == 0 {
				cmd = defaultCmd
			}

			term, err := fac.MakeTtyCmd(cmd)
			if err != nil {
				log.Println(err)
				return
			}

			go func() {
				re := <-resizeCh
				err := term.Resize(re.rows, re.cols)
				if err != nil {
					log.Println(err)
				}
			}()

			opts := []asciitransport.Opt{
				asciitransport.WithReader(term),
				asciitransport.WithWriter(term),
			}
			server.ApplyOpts(opts...)
		}()
	}
}
