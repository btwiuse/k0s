package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"k0s.io/k0s/pkg/asciitransport"
	"k0s.io/k0s/pkg/console"
	"k0s.io/k0s/pkg/localcmd"
	"k0s.io/k0s/pkg/uuid"
)

var (
	A, B = net.Pipe()
)

func server() {
	var (
		factory = &localcmd.Factory{[]string{"bash"}}
		term, _ = factory.New()
		opts    = []asciitransport.Opt{
			asciitransport.WithReader(term),
			asciitransport.WithWriter(term),
		}
		server = asciitransport.Server(B, opts...)
	)

	go func() {
		for {
			var (
				re = <-server.ResizeEvent()
				w  = int(re.Width)
				h  = int(re.Height)
			)
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
		conn net.Conn = A
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
	client := asciitransport.Client(conn, opts...)

	// send
	// r
	go func() {
		sig := make(chan os.Signal)
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
