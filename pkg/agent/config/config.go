package config

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"

	"github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/distro"
	"github.com/btwiuse/conntroll/pkg/name"
	"github.com/btwiuse/conntroll/pkg/uuid"
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
	id   string
	name string
	tags []string
	*url.URL
	verb  bool
	insec bool
	cmd   string
}

func (c *config) Verbose() bool {
	return c.verb
}

func (c *config) Insecure() bool {
	return c.insec
}

func (c *config) Cmd() []string {
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

func (c *config) ID() string {
	return c.id
}

func (c *config) Name() string {
	return c.name
}

func (c *config) Tags() []string {
	return c.tags
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

func (c *config) NewYRPCAgentRequestBody() []byte {
	return []byte(fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nConnection: Keep-Alive\r\n\r\n", "/api/yrpc", c.Hostname()))
}

func (c *config) NewAgentRequestBody() []byte {
	return []byte(fmt.Sprintf("GET %s?%s HTTP/1.1\r\nHost: %s\r\nConnection: Keep-Alive\r\n\r\n", "/api/rpc", c.RawQuery, c.Hostname()))
}

func (c *config) NewSessionRequestBody() []byte {
	return []byte(fmt.Sprintf("GET %s?%s HTTP/1.1\r\nHost: %s\r\nConnection: Keep-Alive\r\n\r\n", "/api/grpc", "id="+c.URL.Query().Get("id"), c.Hostname()))
}

func Parse(args []string) agent.Config {
	var (
		fset        = flag.NewFlagSet("agent", flag.ExitOnError)
		id   string = uuid.New()
		cmd  string

		pwd, _      = os.Getwd()
		_user, _    = user.Current()
		username    = _user.Username
		hostname, _ = os.Hostname()
		goos        = runtime.GOOS
		goarch      = runtime.GOARCH

		hubapi   string = DEFAULT_HUB_ADDRESS
		tags     arrayFlags
		bahash   string
		verbose  bool
		insecure bool

		query url.Values = make(map[string][]string)
		nam              = name.New()
		dist             = distro.Vendor()
	)

	// fset.StringVar(&id, "id", uuid.New(), "Agent ID, for debugging purpose only")
	fset.StringVar(&cmd, "c", "", "Command to run.")
	fset.Var(&tags, "tags", "Agent tags. ")                             // Should be comma separated values like foo,bar
	fset.StringVar(&hubapi, "hub", DEFAULT_HUB_ADDRESS, "Hub address.") //  The 1st positional argument is used if you leave the -hub part.
	fset.BoolVar(&verbose, "verbose", false, "Verbose log.")            //  The 1st positional argument is used if you leave the -hub part.
	fset.BoolVar(&insecure, "insecure", false, "Allow insecure server connections when using SSL")
	fset.StringVar(&bahash, "auth", "conn:troll", "Set user:pass for shell/fs access. Use `-auth offff` to disable auth") // "Protect shell and filesystem access using basicauth. Value should be supplied in user:pass form. (Only hash of user:pass will be sent. hub will use the hash value to authorize access to the agent)"

	err := fset.Parse(args)
	if err != nil {
		log.Println(err)
	}

	if hubapi == DEFAULT_HUB_ADDRESS && len(fset.Args()) != 0 {
		hubapi = fset.Args()[0]
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
	query.Set("tags", tags.String())
	if dist != "" {
		query.Set("distro", dist)
	}

	if strings.Contains(bahash, ":") {
		bahash = fmt.Sprintf("%x", sha256.Sum256([]byte(bahash)))
	} else {
		bahash = ""
	}
	query.Set("bahash", bahash)

	u.RawQuery = query.Encode()

	/* TODO: erase -auth user:pass sensitive info
	defer func(){
		if bahash != "" {
			log.Println(os.Args)
			os.Args[0] = "wtf"
			log.Println(os.Args)
		}
	}()
	*/

	return &config{
		URL:   u,
		id:    id,
		cmd:   cmd,
		name:  nam,
		verb:  verbose,
		insec: insecure,
	}
}
