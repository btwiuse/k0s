package config

import (
	"flag"
	"log"

	"github.com/btwiuse/conntroll/pkg/hub"
)

type config struct {
	port string
	tls  bool
	cert string
	key  string
}

func (c *config) Port() string {
	return c.port
}

func (c *config) UseTLS() bool {
	return c.tls
}

func (c *config) Cert() string {
	return c.cert
}

func (c *config) Key() string {
	return c.key
}

func Parse(args []string) hub.Config {
	fset := flag.NewFlagSet("hub", flag.ExitOnError)
	var (
		port string
		tls  bool
		cert string
		key  string
	)

	fset.StringVar(&port, "l", ":8000", "hub listening port")
	fset.StringVar(&cert, "cert", "", "path to tls cert file")
	fset.StringVar(&key, "key", "", "path to tls key file")
	err := fset.Parse(args)
	if err != nil {
		log.Println(err)
	}

	tls = (cert != "" || key != "")
	if tls {
		if cert == "" {
			log.Fatalln("cert is empty")
		}
		if key == "" {
			log.Fatalln("key is empty")
		}
	}

	return &config{
		port: port,
		tls:  tls,
		cert: cert,
		key:  key,
	}
}
