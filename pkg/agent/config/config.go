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
	return c.URL.Scheme
}

func (c *config) Hostname() string {
	return c.URL.Hostname()
}

func (c *config) NewAgentRequestBody() []byte {
	return []byte(fmt.Sprintf("GET %s?%s HTTP/1.0\r\n\r\n", "/api/rpc", c.RawQuery))
}

func (c *config) NewSessionRequestBody() []byte {
	return []byte(fmt.Sprintf("GET %s?%s HTTP/1.0\r\n\r\n", "/api/session", c.RawQuery))
}

func Parse(args []string) agent.Config {
	fset := flag.NewFlagSet("agent", flag.ExitOnError)
	/*
	fset.Usage = func(){
		fmt.Println("usage:")
	}*/
	var (
		pwd, _      = os.Getwd()
		_user, _    = user.Current()
		username    = _user.Username
		hostname, _ = os.Hostname()
		goos        = runtime.GOOS
		goarch      = runtime.GOARCH

		hubapi string
		// usage bool
		query  url.Values = make(map[string][]string)
	)

	fset.StringVar(&hubapi, "hub", "https://libredot.com", "hub api")
	/*
	fset.BoolVar(&usage, "h", false, "getting help")
	fset.BoolVar(&usage, "help", false, "getting help")
	*/
	err := fset.Parse(args)
	if err != nil {
		log.Println(err)
	}
	/* log.Println(fset.Args())
	if usage {
		fset.Usage()
		os.Exit(0)
	}*/
	if !(strings.HasPrefix(hubapi, "http://") || strings.HasPrefix(hubapi, "https://")) {
		hubapi = "http://" + hubapi
	}

	u, err := url.Parse(hubapi)
	if err != nil {
		log.Fatalln(err)
	}

	id := uuid.New()
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
