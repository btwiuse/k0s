package config

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/btwiuse/conntroll/pkg/uuid"
	"github.com/btwiuse/conntroll/pkg/agent"
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

func (c *config) Addr() string{
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

func (c *config) NewAgentRequestBody() []byte {
	return []byte(fmt.Sprintf("GET %s?%s HTTP/1.0\r\n\r\n", "/api/rpc", c.RawQuery))
}

func (c *config) NewSessionRequestBody() []byte {
	return []byte(fmt.Sprintf("GET %s?%s HTTP/1.0\r\n\r\n", "/api/session", c.RawQuery))
}

func Parse(args []string) agent.Config {
	var (
		fset = flag.NewFlagSet("agent", flag.ExitOnError)
		id string

		pwd, _      = os.Getwd()
		_user, _    = user.Current()
		username    = _user.Username
		hostname, _ = os.Hostname()
		goos        = runtime.GOOS
		goarch      = runtime.GOARCH

		hubapi string

		query  url.Values = make(map[string][]string)
	)

	fset.StringVar(&id, "id", uuid.New(), "agent id, for debugging purpose only")
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
	query.Set("pwd", pwd)
	query.Set("whoami", username)
	query.Set("user", username)
	query.Set("username", username)
	query.Set("hostname", hostname)
	query.Set("os", goos)
	query.Set("arch", goarch)

	u.RawQuery = query.Encode()

	return &config{
		URL: u,
		id:  id,
	}
}
