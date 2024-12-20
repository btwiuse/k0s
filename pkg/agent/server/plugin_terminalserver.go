package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"sync"

	"k0s.io/pkg/agent"
	"k0s.io/pkg/agent/tty/factory"
	"k0s.io/pkg/api"
	"k0s.io/pkg/asciitransport"
)

func init() {
	Tunnels[api.Terminal] = StartTerminalServer
	Channels[api.TerminalID] = StartTerminalServer
}

func StartTerminalServer(c agent.Config) chan net.Conn {
	var (
		ro               bool     = c.GetReadOnly()
		defaultCmd       []string = c.GetCmd()
		terminalListener          = NewLys()
	)
	_ = ro
	go serveTerminal(terminalListener, defaultCmd, c)
	return terminalListener.Conns
}

func serveTerminal(ln net.Listener, defaultCmd []string, c agent.Config) {
	var fac agent.TtyFactory = factory.New(defaultCmd)

	for nth := 1; ; nth++ {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			var (
				tryCommandOnce = &sync.Once{}
				cmdCh          = make(chan []string, 1)
				envCh          = make(chan map[string]string, 1)
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
						envCh <- re.Env
					})
					resizeCh <- struct{ rows, cols int }{rows, cols}
				}
				server.Close()
			}()

			cmd := <-cmdCh
			env := <-envCh

			if len(cmd) == 0 {
				cmd = defaultCmd
			}

			term, err := fac.MakeTtyEnv(cmd, env)
			if err != nil {
				log.Println(err)
				return
			}

			go func() {
				for {
					re := <-resizeCh
					err := term.Resize(re.rows, re.cols)
					if err != nil {
						log.Println(err)
					}
				}
			}()

			logname := fmt.Sprintf("/tmp/%s-%d.log", c.GetID(), nth)
			logfile, err := os.Create(logname)
			if err == nil {
				defer func() {
					exec.Command("dkg-push", logname).Run()
					log.Println("log written to", logname)
				}()
			}

			opts := []asciitransport.Opt{
				asciitransport.WithReader(term),
				asciitransport.WithWriter(term),
				asciitransport.WithLogger(logfile),
			}
			server.ApplyOpts(opts...)

			<-server.Done()
			term.Close()
		}()
	}
}
