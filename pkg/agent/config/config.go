package config

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Config struct {
	HttpServer string
	WsServer   string
	Server     string
	Id         string
}

var Default *Config

func Init() *Config {
	config := new(Config)
	config.Server = "45.32.65.48:8000"
	if len(os.Args) > 1 {
		config.Server = os.Args[1]
	}
	config.HttpServer = fmt.Sprintf("http://%s/tcp-hijack", config.Server)
	config.WsServer = fmt.Sprintf("ws://%s/ws", config.Server)
	config.Id = uuid.New().String()
	Default = config
	return config
}
