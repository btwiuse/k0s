package server

import (
	"net"
	"net/http"

	"github.com/btwiuse/better"
	"k0s.io/pkg/agent/config"
	"k0s.io/pkg/middleware"
)

func StartFileServer(c *config.Config) chan net.Conn {
	var (
		fsListener = NewChannelListener()
		handler    = middleware.LoggingMiddleware(better.FileServer(http.Dir("/")))
		fileServer = &http.Server{Handler: handler}
	)
	go fileServer.Serve(fsListener)
	return fsListener.Conns
}
