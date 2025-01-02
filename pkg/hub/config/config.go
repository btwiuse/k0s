package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/btwiuse/pretty"
	"github.com/btwiuse/version"
	"k0s.io"
	"k0s.io/pkg/utils"
)

type Config struct {
	Port    string
	TLS     bool
	UI      bool
	Verbose bool
	Cert    string
	Key     string
	Ufo     string
	Version *version.Version
}

func Parse(args []string) *Config {
	fset := flag.NewFlagSet("hub", flag.ExitOnError)
	var (
		port        string
		tls         bool
		ui          bool
		verbose     bool
		cert        string
		key         string
		ufo         string
		showVersion bool
	)

	fset.StringVar(&port, "port", utils.EnvPORT(k0s.HUB_PORT), "hub listening port")
	fset.StringVar(&cert, "cert", "", "path to tls cert file")
	fset.StringVar(&key, "key", "", "path to tls key file")
	fset.StringVar(&ufo, "ufo", "", "webteleport station")
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

	c := &Config{
		Port:    port,
		TLS:     tls,
		UI:      ui,
		Cert:    cert,
		Key:     key,
		Ufo:     ufo,
		Verbose: verbose,
		Version: version.Info,
	}

	if showVersion {
		printHubVersion(c)
		os.Exit(0)
	}

	return c
}

type hubVersion struct {
	Hub *version.Version
}

func printHubVersion(c *Config) {
	v := &hubVersion{c.Version}
	fmt.Print(pretty.YAMLString(v))
}
