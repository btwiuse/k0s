package config

import (
	"flag"
	"log"
	"net/url"
	"strings"

	"github.com/btwiuse/conntroll/pkg/client"
	"github.com/btwiuse/conntroll/pkg/uuid"
)

type config struct {
	id string
	*url.URL
}

func (c *config) ID() string {
	return c.id
}

func (c *config) Port() string {
	if c.URL.Port() == "" {
		switch c.URL.Scheme {
		case "http":
			return "80"
		case "https":
			return "443"
		}
	}
	return c.URL.Port()
}

func (c *config) Addr() string {
	return c.Hostname() + ":" + c.Port()
}

func (c *config) Scheme() string {
	if c.URL.Scheme == "http" && c.URL.Hostname() == "" && c.URL.Port() == "443" {
		return "https"
	}
	return c.URL.Scheme
}

func (c *config) Hostname() string {
	if c.URL.Hostname() == "" {
		return "127.0.0.1"
	}
	return c.URL.Hostname()
}

func Parse(args []string) client.Config {
	var (
		fset = flag.NewFlagSet("client", flag.ExitOnError)
		id   string

		hubapi string

		query url.Values = make(map[string][]string)
	)

	fset.StringVar(&id, "id", uuid.New(), "client id, for debugging purpose only")
	fset.StringVar(&hubapi, "hub", "https://libredot.com", "hub api")

	err := fset.Parse(args)
	if err != nil {
		log.Println(err)
	}

	if !(strings.HasPrefix(hubapi, "http://") || strings.HasPrefix(hubapi, "https://")) {
		hubapi = "http://" + hubapi
	}

	u, err := url.Parse(hubapi)
	if err != nil {
		log.Fatalln(err)
	}

	query.Set("id", id)

	u.RawQuery = query.Encode()

	return &config{
		URL: u,
		id:  id,
	}
}
