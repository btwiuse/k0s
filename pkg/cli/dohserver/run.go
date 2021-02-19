package dohserver

import (
	"log"
	"net/http"

	"k0s.io/pkg/dohserver"
	"k0s.io/pkg/tunnel/listener"
)

const exampleConfig = `
remote_addr = "https://chassis.com/dns-query"
path = "/dns-query"
upstream = [
    "udp:1.1.1.1:53",
    "udp:1.0.0.1:53",
    "udp:8.8.8.8:53",
    "udp:8.8.4.4:53",
]
timeout = 10
tries = 3
verbose = true
log_guessed_client_ip = true
ecs_allow_non_global_ip = false
ecs_use_precise_ip = false
`

func Run(args []string) error {
	confPath := "doh-server.conf"
	if len(args) > 0 {
		confPath = args[0]
	}

	log.Println("Loading config from", confPath)
	conf, err := dohserver.LoadConfig(confPath)
	if err != nil {
		log.Println("Example config: " + exampleConfig)
		return err
	}

	server, err := dohserver.NewServer(conf)
	if err != nil {
		return err
	}
	log.Println("Listening on", conf.RemoteAddr)
	return http.Serve(listener.Listener(conf.RemoteAddr, conf.Path), server.Handler())
}
