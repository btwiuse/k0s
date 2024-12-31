package config

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/btwiuse/pretty"
	"github.com/btwiuse/rng"
	"github.com/btwiuse/tags"
	"github.com/btwiuse/version"
	"github.com/denisbrodbeck/machineid"
	"gopkg.in/yaml.v3"

	"k0s.io"
	"k0s.io/pkg/agent/info"
)

type Config struct {
	ID       string                     `json:"id" yaml:"-"`
	Name     string                     `json:"name" yaml:"name"`
	Tags     tags.CommaSeparatedStrings `json:"tags" yaml:"tags"`
	Htpasswd map[string]string          `json:"htpasswd,omitempty" yaml:"htpasswd"`

	Info *info.Info `json:"meta" yaml:"-"`

	Verbose  bool   `json:"-" yaml:"verbose"`
	ReadOnly bool   `json:"-" yaml:"ro"`
	Insecure bool   `json:"-" yaml:"insecure"`
	Pet      bool   `json:"-" yaml:"pet"`
	Cmd      string `json:"-" yaml:"cmd"`
	Hub      string `json:"-" yaml:"hub"`

	uri *url.URL `json:"-"` // where server scheme, host, port, addr are defined

	Version *version.Version `json:"version" yaml:"-"`
}

func (c *Config) GetCmd() []string {
	return c.getCmd()
}

func (c *Config) GetPort() string {
	if c.uri == nil {
		return "443"
	}
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

func (c *Config) GetAddr() string {
	var (
		scheme = c.GetScheme()
		host   = c.GetHost()
		port   = c.GetPort()
	)
	// omit port if already on standard port
	switch {
	case scheme == "http" && port == "80":
		return host
	case scheme == "https" && port == "443":
		return host
	default:
		return fmt.Sprintf("%s:%s", host, port)
	}
}

func (c *Config) GetSchemeWS() string {
	switch c.GetScheme() {
	case "https":
		return "wss"
	default:
		return "ws"
	}
}

func (c *Config) GetScheme() string {
	if c.uri == nil {
		return "https"
	}
	if c.uri.Scheme == "http" && c.uri.Hostname() == "" && c.uri.Port() == "443" {
		return "https"
	}
	return c.uri.Scheme
}

func (c *Config) WithHub(h string) *Config {
	c.Hub = h
	return c
}

func (c *Config) WithCmd(h string) *Config {
	c.Cmd = h
	return c
}

func (c *Config) WithPet(h bool) *Config {
	c.Pet = h
	return c
}

func (c *Config) WithInsecure(h bool) *Config {
	c.Insecure = h
	return c
}

func (c *Config) WithURI() *Config {
	var hubapi = c.Hub
	// default to http
	if !(strings.HasPrefix(hubapi, "http://") || strings.HasPrefix(hubapi, "https://")) {
		hubapi = "http://" + hubapi
	}

	uri, err := url.Parse(hubapi)
	if err != nil {
		log.Fatalln(err)
	}
	c.uri = uri
	return c
}

func (c *Config) WithReadOnly(ro bool) *Config {
	c.ReadOnly = ro
	return c
}

func (c *Config) WithVerbose(v bool) *Config {
	c.Verbose = v
	return c
}

func (c *Config) WithID(id string) *Config {
	c.ID = id
	return c
}

func (c *Config) WithName(name string) *Config {
	c.Name = name
	return c
}

func (c *Config) WithTags(tags []string) *Config {
	c.Tags = append(c.Tags, tags...)
	return c
}

func (c *Config) WithInfo(ifo *info.Info) *Config {
	c.Info = ifo
	return c
}

func (c *Config) GetHost() string {
	if c.uri == nil {
		return "k0s.up.railway.app"
	}
	host := c.uri.Hostname()
	if host == "" {
		return "127.0.0.1"
	}
	return host
}

func isExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func probeConfigFile() string {
	var (
		globalConfig = "/etc/k0s/agent.yaml"
		userConfig   = os.ExpandEnv("${HOME}/.k0s/agent.yaml")
		localConfig  = "agent.yaml"
	)
	for _, conf := range []string{
		localConfig,
		userConfig,
		globalConfig,
	} {
		if isExist(conf) {
			return conf
		}
	}
	return ""
}

func loadConfigFile(file string) *Config {
	c := &Config{
		Hub:     k0s.DEFAULT_HUB_ADDRESS,
		Tags:    []string{},
		Version: version.Info,
	}
	if file == "" {
		return c
	}
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
		return c
	}
	dec := yaml.NewDecoder(f)
	err = dec.Decode(c)
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}
	return c
}

func Parse(args []string) *Config {
	var (
		fset = flag.NewFlagSet("agent", flag.ExitOnError)

		// fset.StringVar(&id, "id", rng.NewUUID(), "Agent ID, for debugging purpose only")
		id string = rng.NewUUID()

		hubapi   *string                    = fset.String("hub", k0s.DEFAULT_HUB_ADDRESS, "Hub address.")
		verbose  *bool                      = fset.Bool("verbose", false, "Verbose log.")
		version  *bool                      = fset.Bool("version", false, "Show agent/hub version info.")
		ro       *bool                      = fset.Bool("ro", false, "Make shell readonly.")
		insecure *bool                      = fset.Bool("insecure", false, "Allow insecure server connections when using SSL.")
		pet      *bool                      = fset.Bool("pet", false, "Run the agent like a pet, on real hardware.")
		name     *string                    = fset.String("name", rng.NewDocker(), "Set agent name.")
		cmd      *string                    = fset.String("cmd", "", "Command to run.")
		c        *string                    = fset.String("c", probeConfigFile(), "Config file location.")
		tags     tags.CommaSeparatedStrings = []string{}
	)

	// Should be comma separated values like foo,bar
	fset.Var(&tags, "tags", "Agent tags.")

	err := fset.Parse(args)
	if err != nil {
		log.Fatalln(err)
	}

	baseConfig := loadConfigFile(*c)

	// Set ID first
	baseConfig.WithID(id)

	// Apply flags if they were set
	fset.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "hub":
			baseConfig.WithHub(*hubapi)
		case "ro":
			baseConfig.WithReadOnly(*ro)
		case "verbose":
			baseConfig.WithVerbose(*verbose)
		case "name":
			baseConfig.WithName(*name)
		case "pet":
			baseConfig.WithPet(*pet)
		case "insecure":
			baseConfig.WithInsecure(*insecure)
		case "tags":
			baseConfig.WithTags(tags)
		case "cmd":
			baseConfig.WithCmd(*cmd)
		}
	})

	// Handle positional hub argument
	if len(fset.Args()) != 0 {
		baseConfig.WithHub(fset.Args()[0])
	}

	// Apply URI and Info
	baseConfig.WithURI().WithInfo(info.CollectInfo())

	// Set default name if not provided
	if baseConfig.Name == "" {
		baseConfig.WithName(*name)
	}

	// Handle pet mode
	if baseConfig.Pet {
		mid, err := machineid.ID()
		if err != nil {
			log.Println(err)
			log.Println("Using alternative approach")
			if mid == "" {
				mid = baseConfig.Info.OS +
					baseConfig.Info.Arch +
					baseConfig.Name +
					baseConfig.Info.Username +
					baseConfig.Info.Hostname
			}
		}
		uid := rng.NewPetUUID(mid)
		baseConfig.WithID(uid)
	}

	if *version {
		printAgentVersion(*baseConfig)
		printHubVersion(*baseConfig)
		os.Exit(0)
	}

	return baseConfig
}

type agentVersion struct {
	Agent *version.Version
}

type hubVersion struct {
	Hub *version.Version
}

func printAgentVersion(c Config) {
	av := &agentVersion{c.Version}
	fmt.Println(pretty.YAMLString(av))
}

func printHubVersion(c Config) {
	var (
		ub = &url.URL{
			Scheme: c.GetScheme(),
			Host:   c.GetAddr(),
			Path:   "/api/version",
		}
		req = &http.Request{
			Method: http.MethodGet,
			URL:    ub,
		}
		t = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: c.Insecure,
				},
			},
		}
	)
	resp, err := t.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// log.Println(string(buf))
	v, err := version.Decode(buf)
	if err != nil {
		log.Fatalln(err)
	}

	hv := &hubVersion{v}
	fmt.Print(pretty.YAMLString(hv))
}

func (c *Config) String() string {
	return pretty.JsonStringLine(c)
}

func Decode(data []byte) (*info.Info, error) {
	v := &Config{
		Info: info.CollectInfo(),
	}
	err := json.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}
	return v.Info, err
}
