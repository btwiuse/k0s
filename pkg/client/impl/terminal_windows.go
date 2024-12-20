package impl

import (
	"log"
	"net"
	"net/url"
	"os"
	"time"

	"github.com/btwiuse/rng"
	"github.com/containerd/console"
	"k0s.io/pkg/asciitransport"
)

func (cl *clientImpl) terminalConnect(endpoint string, userinfo *url.Userinfo) {
	log.Println("Press ESC twice to exit.")

	var (
		c    = cl.Config
		conn net.Conn
		err  error
	)

	for {
		conn, err = cl.Dial(endpoint, userinfo)
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

	opts := []asciitransport.Opt{
		asciitransport.WithReader(os.Stdin),
		asciitransport.WithWriter(os.Stdout),
	}

	if c.GetRecord() {
		logname := rng.NewUUID() + ".log"
		logfile, err := os.Create(logname)
		if err != nil {
			panic(err)
		}
		defer func() {
			log.Println("log written to", logname)
		}()
		opts = append(opts, asciitransport.WithLogger(logfile))
	}

	client := asciitransport.Client(conn, opts...)

	// TODO: find alternative to syscall.SIGWINCH for windows
	// send
	// r
	/*
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
	*/

	<-client.Done()
}
