package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/btwiuse/pretty"
	"k0s.io/conntroll/pkg"
	"k0s.io/conntroll/pkg/hub"
	"k0s.io/conntroll/pkg/version"
)

type config struct {
	port      string
	tls       bool
	cert      string
	key       string
	basicauth bool
	user      string
	pass      string
	version   pkg.Version
}

func (c *config) Port() string {
	return c.port
}

func (c *config) BasicAuth() (string, string, bool) {
	return c.user, c.pass, c.basicauth
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
		basicauth   string
		user        string
		pass        string
		showVersion bool
	)

	fset.StringVar(&port, "port", ":8000", "hub listening port")
	fset.StringVar(&cert, "cert", "", "path to tls cert file")
	fset.StringVar(&key, "key", "", "path to tls key file")
	fset.StringVar(&basicauth, "auth", "", "protect api with basicauth, value should be supplied in user:pass form. Currently only one user:pass pair is supported")
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

	if basicauth != "" {
		ba := strings.SplitN(basicauth, ":", 2)
		user = ba[0]
		pass = ba[1]
	}

	c := &config{
		port:      port,
		tls:       tls,
		cert:      cert,
		key:       key,
		basicauth: basicauth != "",
		user:      user,
		pass:      pass,
		version:   version.GetVersion(),
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
