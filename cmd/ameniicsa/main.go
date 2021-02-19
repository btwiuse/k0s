package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	asciitransport "k0s.io/pkg/asciitransport/v2"
	"k0s.io/pkg/console"
	"k0s.io/pkg/localcmd"
	"k0s.io/pkg/uuid"
)

var (
	Client, Server = net.Pipe()
)

func server() {
	var (
		factory = &localcmd.Factory{[]string{"bash"}}
		term, _ = factory.New()
		opts    = []asciitransport.Opt{
			asciitransport.WithReader(term),
			asciitransport.WithWriter(term),
			asciitransport.WithResizeHook(func(w, h uint16){
				err := term.Resize(int(w), int(h))
				if err != nil {
					log.Println(err)
				}
			}),
		}
		server = asciitransport.Server(Server, opts...)
	)

	<-server.Done()
	log.Println("detached", term.Close())
}

func main() {
	log.Println("Press ESC twice to exit.")

	go server()

	var (
		err  error
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
				uint16(currentSize.Width)-1,
				uint16(currentSize.Height)-1,
			)

			switch <-sig {
			case syscall.SIGWINCH:
			}
		}
	}()

	<-client.Done()
}
