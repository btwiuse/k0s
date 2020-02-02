package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/btwiuse/pretty"
	"k0s.io/k0s/pkg"
	"k0s.io/k0s/pkg/hub"
	"k0s.io/k0s/pkg/version"
)

type config struct {
	port    string
	tls     bool
	cert    string
	key     string
	version pkg.Version
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

func (c *config) GetVersion() pkg.Version {
	return c.version
}

func Parse(args []string) hub.Config {
	fset := flag.NewFlagSet("hub", flag.ExitOnError)
	var (
		port        string
		tls         bool
		cert        string
		key         string
		showVersion bool
	)

	fset.StringVar(&port, "port", ":8000", "hub listening port")
	fset.StringVar(&cert, "cert", "", "path to tls cert file")
	fset.StringVar(&key, "key", "", "path to tls key file")
	fset.BoolVar(&showVersion, "version", false, "Show hub version info.")
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

	c := &config{
		port:    port,
		tls:     tls,
		cert:    cert,
		key:     key,
		version: version.GetVersion(),
	}

	if showVersion {
		printHubVersion(c)
		os.Exit(0)
	}

	return c
}

type hubVersion struct {
	Hub pkg.Version
}

func printHubVersion(c hub.Config) {
	v := &hubVersion{c.GetVersion()}
	fmt.Print(pretty.YAMLString(v))
}
