package config

import (
	"flag"
	"log"

	"github.com/btwiuse/conntroll/pkg/hub"
)

type config struct {
	port string
}

func (c *config) Port() string {
	return c.port
}

func Parse(args []string) hub.Config {
	fset := flag.NewFlagSet("hub", flag.ExitOnError)
	var (
		port string
	)

	fset.StringVar(&port, "l", ":8000", "hub listening port")
	err := fset.Parse(args)
	if err != nil {
		log.Println(err)
	}

	return &config{
		port: port,
	}
}
