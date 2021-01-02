package dohserver

import (
	"log"
	"net/http"

	"k0s.io/k0s/pkg/tunnel/listener"
)

func Run(args []string) error {
	confPath := "doh-server.conf"
	if len(args) > 0 {
		confPath = args[0]
	}

	log.Println("Loading config from", confPath)
	conf, err := loadConfig(confPath)
	if err != nil {
		log.Println("Example config: " + exampleConfig)
		return err
	}

	server, err := NewServer(conf)
	if err != nil {
		return err
	}
	log.Println("Listening on", conf.RemoteAddr)
	return http.Serve(listener.Listener(conf.RemoteAddr, conf.Path), server.Handler())
}
