// +build linux darwin
// +build !windows

package client

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

	"github.com/btwiuse/asciitransport"
	"github.com/btwiuse/wetty/pkg/utils"
	"github.com/containerd/console"
	"github.com/gorilla/websocket"
	"k0s.io/k0s/pkg/uuid"
)

func dial(p string, h http.Header) (conn net.Conn, err error) {
	wsconn, _, err := websocket.DefaultDialer.Dial(p, h)
	if err != nil {
		return nil, err
	}

	return utils.NetConn(wsconn), nil
}

func terminalConnect(endpoint string, userinfo *url.Userinfo) {
	log.Println("Press ESC twice to exit.")

	var (
		conn net.Conn
		err  error
		h    = http.Header{
			"Authorization": {
				"Basic " + base64.StdEncoding.EncodeToString([]byte(userinfo.String())),
			},
		}
	)
	for {
		// conn, err = net.Dial("tcp", ":12345")
		conn, err = dial(endpoint, h)
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
