package config

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"strings"

	"github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/agent/info"
	"github.com/btwiuse/conntroll/pkg/rng"
	"github.com/btwiuse/conntroll/pkg/uuid"
	"github.com/btwiuse/pretty"
)

const DEFAULT_HUB_ADDRESS = "https://hub.libredot.com"

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ",")
}

func (i *arrayFlags) Set(value string) error {
	if value == "" {
		return nil
	}
	tags := strings.Split(value, ",")
	*i = append(*i, tags...)
	return nil
}

type config struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
	Auth string   `json:"auth,omitempty"`

	agent.Info `json:"meta"`

	verbose  bool
	insecure bool
	cmd      string
	uri      *url.URL `json:"-"` // where server scheme, host, port, addr are defined
}

func (c *config) GetVerbose() bool {
	return c.verbose
}

func (c *config) GetInsecure() bool {
	return c.insecure
}

func (c *config) GetCmd() []string {
	shell := "bash"
	if _, err := exec.LookPath(shell); err != nil {
		shell = "sh"
	}
	args := []string{"/usr/bin/env", "TERM=xterm", shell}
	if c.cmd == "" {
		return args
	}
	return append(args, "-c", c.cmd)
}

func (c *config) GetID() string {
	return c.ID
}

func (c *config) GetName() string {
	return c.Name
}

func (c *config) GetTags() []string {
	return c.Tags
}

func (c *config) GetPort() string {
	if c.uri.Port() == "" {
		switch c.uri.Scheme {
		case "http":
			return "80"
		case "https":
			return "443"
		}
	}
	return c.uri.Port()
}

func (c *config) GetAddr() string {
	return c.GetHost() + ":" + c.GetPort()
}

func (c *config) GetScheme() string {
	if c.uri.Scheme == "http" && c.uri.Hostname() == "" && c.uri.Port() == "443" {
		return "https"
	}
	return c.uri.Scheme
}

func (c *config) GetHost() string {
	host := c.uri.Hostname()
	if host == "" {
		return "127.0.0.1"
	}
	return host
}

func (c *config) FakeHeader(p string) []byte {
	return []byte(
		fmt.Sprintf(
			"GET %s HTTP/1.1\r\nHost: %s\r\nConnection: Keep-Alive\r\n\r\n",
			p,
			c.GetHost(),
		),
	)
}

func Parse(args []string) agent.Config {
	var (
		fset = flag.NewFlagSet("agent", flag.ExitOnError)

		id              = uuid.New()
		name            = rng.New()
		tags arrayFlags = []string{}

		hubapi   string = DEFAULT_HUB_ADDRESS
		bahash   string
		verbose  bool
		insecure bool
		cmd      string
	)

	// fset.StringVar(&id, "id", uuid.New(), "Agent ID, for debugging purpose only")

	// Should be comma separated values like foo,bar
	fset.Var(&tags, "tags", "Agent tags. ")

	// "Protect shell and filesystem access using basicauth. Value should be supplied in user:pass form.
	// (Only hash of user:pass will be sent. hub will use the hash value to authorize access to the agent)"
	fset.StringVar(&bahash, "auth", "conn:troll", "Set user:pass for shell/fs access. Use `-auth offff` to disable auth")

	//  The 1st positional argument is used if you leave the -hub part.
	fset.StringVar(&hubapi, "hub", DEFAULT_HUB_ADDRESS, "Hub address.")

	//  The 1st positional argument is used if you leave the -hub part.
	fset.BoolVar(&verbose, "verbose", false, "Verbose log.")

	fset.BoolVar(&insecure, "insecure", false, "Allow insecure server connections when using SSL")

	fset.StringVar(&cmd, "c", "", "Command to run.")

	err := fset.Parse(args)
	if err != nil {
		log.Fatalln(err)
	}

	if hubapi == DEFAULT_HUB_ADDRESS && len(fset.Args()) != 0 {
		hubapi = fset.Args()[0]
	}

	// default to http
	if !(strings.HasPrefix(hubapi, "http://") || strings.HasPrefix(hubapi, "https://")) {
		hubapi = "http://" + hubapi
	}

	if strings.Contains(bahash, ":") {
		bahash = fmt.Sprintf("%x", sha256.Sum256([]byte(bahash)))
	} else {
		bahash = ""
	}

	uri, err := url.Parse(hubapi)
	if err != nil {
		log.Fatalln(err)
	}

	return &config{
		Info: info.CollectInfo(),

		uri: uri,

		ID:   id,
		Name: name,
		Tags: tags,
		Auth: bahash,

		verbose:  verbose,
		insecure: insecure,
		cmd:      cmd,
	}
}

func (c *config) String() string {
	return strings.TrimSpace(pretty.JSONString(c))
}

func Decode(data []byte) (agent.Info, error) {
	v := &config{
		Info: info.EmptyInfo(),
	}
	err := json.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}
	return v, err
}
