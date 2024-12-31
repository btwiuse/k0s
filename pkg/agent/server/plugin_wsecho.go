package server

import (
	"io"
	"log"
	"net"
	"net/http"

	"github.com/btwiuse/wsconn"
	"k0s.io/pkg/agent/config"
	"k0s.io/pkg/middleware"
)

func StartWsEchoServer(c *config.Config) chan net.Conn {
	var (
		wsEchoListener = NewChannelListener()
		handler        = middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Header)
			conn, err := wsconn.Wrconn(w, r)
			if err != nil {
				log.Println(err)
				return
			}

			io.Copy(conn, conn)
		}))
		wsEchoServer = &http.Server{Handler: handler}
	)
	go wsEchoServer.Serve(wsEchoListener)
	return wsEchoListener.Conns
}

func StartWsEcho2Server(c *config.Config) chan net.Conn {
	var (
		wsEchoListener = NewWSChannelListener()
	)
	go handleWsEcho2(wsEchoListener)
	return wsEchoListener.Conns
}

func handleWsEcho2(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			defer conn.Close()
			io.Copy(conn, conn)
		}()
	}
}

func StartWsEcho3Server(c *config.Config) chan net.Conn {
	var (
		wsEchoListener = NewChannelListener()
	)
	go handleWsEcho2(wsEchoListener)
	return wsEchoListener.Conns
}
