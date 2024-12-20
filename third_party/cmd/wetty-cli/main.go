package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/btwiuse/rng"
	"github.com/containerd/console"
	"k0s.io/pkg/utils"
	asciitransport "k0s.io/third_party/pkg/asciitransport/v2"
	"k0s.io/third_party/pkg/wetty/wetty"
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

	addr := "ws://127.0.0.1:45080/terminal"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	log.Println(`addr = os.Args[1] || "ws://127.0.0.1:45080/terminal"`)

	var (
		conn net.Conn
		err  error
	)
	for {
		log.Println("Connecting to", addr)
		conn, err = dial(addr)
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
	}
	client := asciitransport.Client(conn, opts...)

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
				uint16(currentSize.Height),
				uint16(currentSize.Width),
			)

			switch <-sig {
			case syscall.SIGWINCH:
			}
		}
	}()

	<-client.Done()
}
