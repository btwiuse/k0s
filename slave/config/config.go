package config

import (
	"fmt"
	"os"
)

type Config struct {
	WsServer string
	Server   string
	Id       string
}

var Default *Config

func Init() *Config {
	server := "45.32.65.48:8000"
	if len(os.Args) > 1 {
		server = os.Args[1]
	}
	return &Config{
		Server:   server,
		WsServer: fmt.Sprintf("ws://%s/ws", server),
	}
}
