package config

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/name"
	"github.com/btwiuse/conntroll/pkg/uuid"
)

type config struct {
	id   string
	name string
	*url.URL
}

func (c *config) ID() string {
	return c.id
}

func (c *config) Name() string {
	return c.name
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

func (c *config) NewAgentRequestBody() []byte {
	return []byte(fmt.Sprintf("GET %s?%s HTTP/1.1\r\nHost: %s\r\nConnection: Keep-Alive\r\n\r\n", "/api/rpc", c.RawQuery, c.Hostname()))
}

func (c *config) NewSessionRequestBody() []byte {
	return []byte(fmt.Sprintf("GET %s?%s HTTP/1.1\r\nHost: %s\r\nConnection: Keep-Alive\r\n\r\n", "/api/grpc", "id="+c.URL.Query().Get("id"), c.Hostname()))
}

func Parse(args []string) agent.Config {
	var (
		fset = flag.NewFlagSet("agent", flag.ExitOnError)
		id   string

		pwd, _      = os.Getwd()
		_user, _    = user.Current()
		username    = _user.Username
		hostname, _ = os.Hostname()
		goos        = runtime.GOOS
		goarch      = runtime.GOARCH

		hubapi string
		bahash string

		query url.Values = make(map[string][]string)
		nam = name.New()
	)

	fset.StringVar(&id, "id", uuid.New(), "agent id, for debugging purpose only")
	fset.StringVar(&hubapi, "hub", "https://libredot.com", "hub api")
	fset.StringVar(&bahash, "basicauth", "", "protect api with basicauth, value should be supplied in user:pass form. (Only hash of user:pass will be sent. hub will use the hash value to authorize access to the agent)")

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
	query.Set("username", username)
	query.Set("hostname", hostname)
	query.Set("os", goos)
	query.Set("arch", goarch)
	query.Set("name", nam)

	if bahash != "" {
		bahash = fmt.Sprintf("%x", sha256.Sum256([]byte(bahash)))
	}
	query.Set("bahash", bahash)

	u.RawQuery = query.Encode()

	return &config{
		URL: u,
		id:  id,
		name:  nam,
	}
}
