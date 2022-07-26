package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/btwiuse/pretty"
	"k0s.io"
	"k0s.io/pkg/hub"
	"k0s.io/pkg/version"
)

type config struct {
	port    string
	tls     bool
	ui      bool
	verbose bool
	cert    string
	key     string
	version k0s.Version
}

func (c *config) Port() string {
	if c.port != "" {
		return c.port
	}
	if port, ok := os.LookupEnv("PORT"); ok {
		return ":" + port
	}
	return ":8000"
}

func (c *config) UseTLS() bool {
	return c.tls
}

func (c *config) Verbose() bool {
	return c.verbose
}

func (c *config) UI() bool {
	return c.ui
}

func (c *config) Cert() string {
	return c.cert
}

func (c *config) Key() string {
	return c.key
}

func (c *config) GetVersion() k0s.Version {
	return c.version
}

func Parse(args []string) hub.Config {
	fset := flag.NewFlagSet("hub", flag.ExitOnError)
	var (
		port        string
		tls         bool
		ui          bool
		verbose     bool
		cert        string
		key         string
		showVersion bool
	)

	fset.StringVar(&port, "port", "", "hub listening port")
	fset.StringVar(&cert, "cert", "", "path to tls cert file")
	fset.StringVar(&key, "key", "", "path to tls key file")
	fset.BoolVar(&showVersion, "version", false, "Show hub version info.")
	fset.BoolVar(&ui, "ui", false, "enable web ui.")
	fset.BoolVar(&verbose, "verbose", false, "show verbose log.")
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
		ui:      ui,
		cert:    cert,
		key:     key,
		verbose: verbose,
		version: version.GetVersion(),
	}

	if showVersion {
		printHubVersion(c)
		os.Exit(0)
	}

	return c
}

type hubVersion struct {
	Hub k0s.Version
}

func printHubVersion(c hub.Config) {
	v := &hubVersion{c.GetVersion()}
	fmt.Print(pretty.YAMLString(v))
}
