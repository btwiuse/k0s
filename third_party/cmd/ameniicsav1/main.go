package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/btwiuse/rng"
	"github.com/containerd/console"
	"k0s.io/pkg/agent/tty/factory"
	"k0s.io/pkg/asciitransport"
)

var (
	Client, Server = net.Pipe()
)

func server() {
	var (
		fac     = factory.New([]string{"bash"})
		term, _ = fac.MakeTty()
		session = asciitransport.Server(Server)
	)

	// Handle resize events
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
	log.Println("detached", term.Close())
}

func main() {
	log.Println("Press ESC twice to exit.")

	go server()

	var (
		err error
	)

	term, err := console.ConsoleFromFile(os.Stdin)
	if err != nil {
		panic(err)
	}
	defer term.Reset()

	if err = term.SetRaw(); err != nil {
		panic(err)
	}

	logname := rng.NewUUID() + ".log"
	logfile, err := os.Create(logname)
	if err != nil {
		panic(err)
	}
	defer func() {
		log.Println("log written to", logname)
	}()

	opts := []asciitransport.Opt{
		asciitransport.WithLogger(logfile),
		asciitransport.WithReader(os.Stdin),
		asciitransport.WithWriter(os.Stdout),
		asciitransport.WithCommand([]string{"neofetch"}),
		asciitransport.WithEnv(map[string]string{
			"NEW": "Zealand",
		}),
	}
	client := asciitransport.Client(Client, opts...)

	// send
	// r
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGWINCH)

		for {
			currentSize, err := term.Size()
			if err != nil {
				log.Println(err)
				continue
			}

			// log.Println(currentSize)
			client.Resize(
				uint(currentSize.Width),
				uint(currentSize.Height),
			)

			switch <-sig {
			case syscall.SIGWINCH:
			}
		}
	}()

	<-client.Done()
}
