package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/containerd/console"
	"k0s.io/pkg/agent/tty/factory"
	"k0s.io/pkg/asciitransport"
	"k0s.io/pkg/uuid"
)

var (
	Client, Server = net.Pipe()
)

func server() {
	var (
		fac     = factory.New([]string{"bash"})
		term, _ = fac.MakeTty()
		opts    = []asciitransport.Opt{
			asciitransport.WithReader(term),
			asciitransport.WithWriter(term),
			/*
				asciitransport.WithResizeHook(func(w, h uint16){
					err := term.Resize(int(w), int(h))
					if err != nil {
						log.Println(err)
					}
				}),
			*/
		}
		server = asciitransport.Server(Server, opts...)
	)

	go func() {
		for {
			var (
				re = <-server.ResizeEvent()
				w  = int(re.Width)
				h  = int(re.Height)
			)
			_ = w
			_ = h
			err := term.Resize(w, h)
			if err != nil {
				log.Println(err)
				break
			}
		}
		server.Close()
	}()

	<-server.Done()
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

	logname := uuid.New() + ".log"
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
