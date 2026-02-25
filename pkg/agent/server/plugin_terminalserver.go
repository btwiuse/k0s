package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"

	"k0s.io/pkg/agent"
	"k0s.io/pkg/agent/config"
	"k0s.io/pkg/agent/tty/factory"
	"k0s.io/pkg/asciitransport"
)

func StartTerminalServer(c *config.Config) chan net.Conn {
	var (
		ro               bool     = c.ReadOnly
		defaultCmd       []string = c.GetCmd()
		terminalListener          = NewChannelListener()
	)
	_ = ro
	go serveTerminal(terminalListener, defaultCmd, c)
	return terminalListener.Conns
}

func serveTerminal(ln net.Listener, defaultCmd []string, c *config.Config) {
	var fac agent.TtyFactory = factory.New(defaultCmd)

	for nth := 1; ; nth++ {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			logname := fmt.Sprintf("/tmp/%s-%d.log", c.ID, nth)
			logfile, err := os.Create(logname)

			var opts []asciitransport.Opt
			if err == nil {
				defer func() {
					exec.Command("dkg-push", logname).Run()
					log.Println("log written to", logname)
				}()
				opts = append(opts, asciitransport.WithLogger(logfile))
			}

			session := asciitransport.Server(conn, opts...)

			// First resize carries command and env
			re, err := session.NextResize()
			if err != nil {
				log.Println(err)
				return
			}

			cmd := re.Command
			env := re.Env

			if len(cmd) == 0 {
				cmd = defaultCmd
			}

			term, err := fac.MakeTtyEnv(cmd, env)
			if err != nil {
				log.Println(err)
				return
			}

			// Apply initial size
			if resizeErr := term.Resize(int(re.Width), int(re.Height)); resizeErr != nil {
				log.Println(resizeErr)
			}

			// Handle future resize events
			go func() {
				for {
					re, err := session.NextResize()
					if err != nil {
						break
					}
					if resizeErr := term.Resize(int(re.Width), int(re.Height)); resizeErr != nil {
						log.Println(resizeErr)
					}
				}
			}()

			// Bridge I/O between session and PTY
			done := make(chan struct{})
			go func() {
				io.Copy(term, session)
				close(done)
			}()
			go io.Copy(session, term)

			<-done
			session.Close()
			term.Close()
		}()
	}
}
