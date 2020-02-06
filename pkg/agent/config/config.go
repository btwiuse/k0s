package config

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/btwiuse/pretty"
	"github.com/denisbrodbeck/machineid"
	"gopkg.in/yaml.v2"
	"k0s.io/k0s/pkg"
	"k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/agent/info"
	"k0s.io/k0s/pkg/rng"
	"k0s.io/k0s/pkg/uuid"
	"k0s.io/k0s/pkg/version"
)

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
	ID       string            `json:"id" yaml:"-"`
	Name     string            `json:"name" yaml:"name"`
	Tags     []string          `json:"tags" yaml:"tags"`
	Htpasswd map[string]string `json:"htpasswd,omitempty" yaml:"htpasswd"`

	agent.Info `json:"meta" yaml:"-"`

	Verbose  bool   `json:"-" yaml:"verbose"`
	ReadOnly bool   `json:"-" yaml:"ro"`
	Insecure bool   `json:"-" yaml:"insecure"`
	Pet      bool   `json:"-" yaml:"pet"`
	Cmd      string `json:"-" yaml:"cmd"`
	Hub      string `json:"-" yaml:"hub"`

	uri *url.URL `json:"-"` // where server scheme, host, port, addr are defined

	Version pkg.Version `json:"version" yaml:"-"`
}

func (c *config) GetVersion() pkg.Version {
	return c.Version
}

func (c *config) GetVerbose() bool {
	return c.Verbose
}

func (c *config) GetReadOnly() bool {
	return c.ReadOnly
}

func (c *config) GetInsecure() bool {
	return c.Insecure
}

func (c *config) GetPet() bool {
	return c.Pet
}

func (c *config) GetCmd() []string {
	return c.getCmd()
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

func (c *config) GetSchemeWS() string {
	switch c.GetScheme() {
	case "https":
		return "wss"
	default:
		return "ws"
	}
}

func (c *config) GetScheme() string {
	if c.uri.Scheme == "http" && c.uri.Hostname() == "" && c.uri.Port() == "443" {
		return "https"
	}
	return c.uri.Scheme
}

type Opt func(c *config)

func SetHub(h string) Opt {
	return func(c *config) {
		c.Hub = h
	}
}

func SetCmd(h string) Opt {
	return func(c *config) {
		c.Cmd = h
	}
}

func SetPet(h bool) Opt {
	return func(c *config) {
		c.Pet = h
	}
}

func SetInsecure(h bool) Opt {
	return func(c *config) {
		c.Insecure = h
	}
}

func SetURI() Opt {
	return func(c *config) {
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
	}
}

func SetReadOnly(ro bool) Opt {
	return func(c *config) {
		c.ReadOnly = ro
	}
}

func SetVerbose(v bool) Opt {
	return func(c *config) {
		c.Verbose = v
	}
}

func SetID(id string) Opt {
	return func(c *config) {
		c.ID = id
	}
}

func SetName(name string) Opt {
	return func(c *config) {
		c.Name = name
	}
}

func SetTags(tags []string) Opt {
	return func(c *config) {
		c.Tags = append(c.Tags, tags...)
	}
}

func SetInfo(ifo agent.Info) Opt {
	return func(c *config) {
		c.Info = ifo
	}
}

func (c *config) GetHost() string {
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

func loadConfigFile(file string) *config {
	c := &config{
		Hub:     pkg.DEFAULT_HUB_ADDRESS,
		Tags:    []string{},
		Version: version.GetVersion(),
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

func Parse(args []string) agent.Config {
	var (
		fset = flag.NewFlagSet("agent", flag.ExitOnError)

		// fset.StringVar(&id, "id", uuid.New(), "Agent ID, for debugging purpose only")
		id = uuid.New()

		opts = []Opt{
			SetID(id),
		}

		hubapi   *string    = fset.String("hub", pkg.DEFAULT_HUB_ADDRESS, "Hub address.")
		verbose  *bool      = fset.Bool("verbose", false, "Verbose log.")
		version  *bool      = fset.Bool("version", false, "Show agent/hub version info.")
		ro       *bool      = fset.Bool("ro", false, "Make shell readonly.")
		insecure *bool      = fset.Bool("insecure", false, "Allow insecure server connections when using SSL.")
		pet      *bool      = fset.Bool("pet", false, "Run the agent like a pet, on real hardware.")
		name     *string    = fset.String("name", rng.New(), "Set agent name.")
		cmd      *string    = fset.String("cmd", "", "Command to run.")
		c        *string    = fset.String("c", probeConfigFile(), "Config file location.")
		tags     arrayFlags = []string{}
	)

	// Should be comma separated values like foo,bar
	fset.Var(&tags, "tags", "Agent tags.")

	err := fset.Parse(args)
	if err != nil {
		log.Fatalln(err)
	}

	fset.Visit(func(f *flag.Flag) {
		if f.Name == "hub" {
			opts = append(opts, SetHub(*hubapi))
		}
		if f.Name == "ro" {
			opts = append(opts, SetReadOnly(*ro))
		}
		if f.Name == "verbose" {
			opts = append(opts, SetVerbose(*verbose))
		}
		if f.Name == "name" {
			opts = append(opts, SetName(*name))
		}
		if f.Name == "pet" {
			opts = append(opts, SetPet(*pet))
		}
		if f.Name == "insecure" {
			opts = append(opts, SetInsecure(*insecure))
		}
		if f.Name == "tags" {
			opts = append(opts, SetTags(tags))
		}
		if f.Name == "cmd" {
			opts = append(opts, SetCmd(*cmd))
		}
	})

	//  The 1st positional argument is used if you leave out the -hub part.
	if len(fset.Args()) != 0 {
		opts = append(opts, SetHub(fset.Args()[0]))
	}

	opts = append(opts, SetURI())
	opts = append(opts, SetInfo(info.CollectInfo()))

	baseConfig := loadConfigFile(*c)

	if baseConfig.GetName() == "" {
		opts = append(opts, SetName(*name))
	}

	for _, opt := range opts {
		opt(baseConfig)
	}

	if baseConfig.GetPet() {
		mid, err := machineid.ID()
		if err != nil {
			log.Println(err)
			log.Println("Using alternative approach")
			// on some platforms like android, mid is empty string
			// assume user has set a fixed name
			// generate a fixed id with best effort
			// based on provided info
			// use mid as seed
			if mid == "" {
				mid = baseConfig.GetOS() +
					baseConfig.GetArch() +
					baseConfig.GetName() +
					baseConfig.GetUsername() +
					baseConfig.GetHostname()
			}
		}
		uid := uuid.NewPet(mid)
		SetID(uid)(baseConfig)
	}

	if *version {
		printAgentVersion(baseConfig)
		printHubVersion(baseConfig)
		os.Exit(0)
	}

	return baseConfig
}

type agentVersion struct {
	Agent pkg.Version
}

type hubVersion struct {
	Hub pkg.Version
}

func printAgentVersion(c agent.Config) {
	av := &agentVersion{c.GetVersion()}
	fmt.Println(pretty.YAMLString(av))
}

func printHubVersion(c agent.Config) {
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
					InsecureSkipVerify: c.GetInsecure(),
				},
			},
		}
	)
	resp, err := t.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	buf, err := ioutil.ReadAll(resp.Body)
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

func (c *config) String() string {
	return pretty.JsonString(c)
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
