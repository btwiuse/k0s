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
)

type Config struct {
	ID string
	*url.URL
}

func (c *Config) NewAgentRequestBody() []byte {
	return []byte(fmt.Sprintf("GET %s?%s HTTP/1.0\r\n\r\n", "/api/rpc", c.RawQuery))
}

func (c *Config) NewSessionRequestBody() []byte {
	return []byte(fmt.Sprintf("GET %s?%s HTTP/1.0\r\n\r\n", "/api/session", c.RawQuery))
}

func Parse() *Config {
	var (
		pwd, _      = os.Getwd()
		_user, _    = user.Current()
		username    = _user.Username
		hostname, _ = os.Hostname()
		goos        = runtime.GOOS
		goarch      = runtime.GOARCH

		hubapi string
		query  url.Values = make(map[string][]string)
	)

	flag.StringVar(&hubapi, "hub", "https://libredot.com", "hub api")
	flag.Parse()
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

	return &Config{
		URL: u,
		ID:  id,
	}
}
