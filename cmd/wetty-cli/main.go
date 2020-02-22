package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"k0s.io/k0s/pkg/asciitransport"
	"k0s.io/k0s/pkg/console"
	"k0s.io/k0s/pkg/utils"
	"k0s.io/k0s/pkg/uuid"
	"k0s.io/k0s/pkg/wetty/wetty"
)

func dial(p string) (conn net.Conn, err error) {
	dialer := wetty.Dialer
	wsconn, _, err := dialer.Dial(p, nil)
	if err != nil {
		return nil, err
	}

	return utils.NetConn(wsconn), nil
}

func main() {
	log.Println("Press ESC twice to exit.")

	var (
		conn net.Conn
		err  error
	)
	for {
		// conn, err = net.Dial("tcp", ":12345")
		conn, err = dial("ws://127.0.0.1:8080/terminal")
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}
		break
	}

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
				uint(currentSize.Height),
				uint(currentSize.Width),
			)

			switch <-sig {
			case syscall.SIGWINCH:
			}
		}
	}()

	<-client.Done()
}
