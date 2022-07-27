//go:build !windows
// +build !windows

package impl

import (
	"encoding/base64"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"k0s.io/pkg/asciitransport"
	"k0s.io/pkg/client/dial"
	"k0s.io/pkg/console"
	"k0s.io/pkg/uuid"
)

func (cl *clientImpl) terminalConnect(endpoint string, userinfo *url.Userinfo) {
	log.Println("Press ESC twice to exit. (Some terminals don't supported it)")

	var (
		c    = cl.Config
		conn net.Conn
		err  error
		h    = http.Header{
			"Authorization": {
				"Basic " + base64.StdEncoding.EncodeToString([]byte(userinfo.String())),
			},
		}
		dialer = dial.New(c)
	)

	for {
		conn, err = dialer.Dial(endpoint, h)
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
		logname := uuid.New() + ".log"
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
