//go:build plugin_dohserver

package server

import (
	"log"
	"net"
	"net/http"

	"k0s.io/pkg/agent"
	"k0s.io/pkg/api"
	"k0s.io/pkg/middleware"
	"k0s.io/third_party/pkg/dohserver"
)

func init() { Tunnels[api.Doh] = StartDohServer }

func newDohHandler() (http.Handler, error) {
	confPath := "doh-server.conf"
	log.Println("dohserver: Loading config from", confPath)
	conf, err := dohserver.LoadConfig(confPath)
	if err != nil {
		return nil, err
	}

	server, err := dohserver.NewServer(conf)
	if err != nil {
		return nil, err
	}
	return server.Handler(), nil
}

func StartDohServer(c agent.Config) chan net.Conn {
	dohHandler, err := newDohHandler()
	if err != nil {
		log.Println("dohserver: initialization failed:", err)
		dohHandler = http.NotFoundHandler()
	}
	var (
		dohListener = NewChannelListener()
		handler     = middleware.LoggingMiddleware(dohHandler)
		dohServer   = &http.Server{Handler: handler}
	)
	go dohServer.Serve(dohListener)
	return dohListener.Conns
}
