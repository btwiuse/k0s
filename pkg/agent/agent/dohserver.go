package agent

import (
	"log"
	"net"
	"net/http"

	types "k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/dohserver"
	"k0s.io/k0s/pkg/middleware"
)

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

func StartDohServer(c types.Config) chan net.Conn {
	dohHandler, err := newDohHandler()
	if err != nil {
		log.Println("dohserver: initialization failed:", err)
		dohHandler = http.NotFoundHandler()
	}
	var (
		dohListener = NewLys()
		handler     = middleware.LoggingMiddleware(dohHandler)
		dohServer   = &http.Server{Handler: handler}
	)
	go dohServer.Serve(dohListener)
	return dohListener.Conns
}
