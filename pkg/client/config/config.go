package config

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/btwiuse/pretty"
	"gopkg.in/yaml.v3"

	"github.com/btwiuse/version"
	"k0s.io"
)

type KeyStore map[string]string

type Config struct {
	Redir            string              `json:"-" yaml:"redir"`
	Socks            string              `json:"-" yaml:"socks"`
	Doh              string              `json:"-" yaml:"doh"`
	Verbose          bool                `json:"-" yaml:"verbose"`
	Insecure         bool                `json:"-" yaml:"insecure"`
	Record           bool                `json:"-" yaml:"record"`
	CacheCredentials bool                `json:"-" yaml:"cache_credentials"`
	Credentials      map[string]KeyStore `json:"-" yaml:"credentials"`
	ConfigLocation   string              `json:"-" yaml:"-"`
	Hub              string              `json:"-" yaml:"hub"`

	uri *url.URL `json:"-"` // where server scheme, host, port, addr are defined

	Version *version.Version `json:"version" yaml:"-"`
}

func (c *Config) GetPort() string {
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
	if c.uri.Scheme == "http" && c.uri.Hostname() == "" && c.uri.Port() == "443" {
		return "https"
	}
	return c.uri.Scheme
}

func (c *Config) GetHost() string {
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
		globalConfig = "/etc/k0s/client.yaml"
		userConfig   = os.ExpandEnv("${HOME}/.k0s/client.yaml")
		localConfig  = "client.yaml"
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
		Hub:            k0s.DEFAULT_HUB_ADDRESS,
		Version:        version.Info,
		ConfigLocation: file,
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
		fset = flag.NewFlagSet("client", flag.ExitOnError)

		hubapi   *string = fset.String("hub", k0s.DEFAULT_HUB_ADDRESS, "Hub address.")
		redir    *string = fset.String("redir", k0s.REDIR_PROXY_PORT, "Redir port.")
		socks    *string = fset.String("socks", k0s.SOCKS5_PROXY_PORT, "Socks port.")
		doh      *string = fset.String("doh", k0s.DOH_PROXY_PORT, "Doh port.")
		verbose  *bool   = fset.Bool("verbose", false, "Verbose log.")
		version  *bool   = fset.Bool("version", false, "Show agent/hub version info.")
		insecure *bool   = fset.Bool("insecure", false, "Allow insecure server connections when using SSL.")
		record   *bool   = fset.Bool("record", false, "Record terminal events to a log file.")
		// cc       *bool   = fset.Bool("cc", false, "Cache credentials.")
		c *string = fset.String("c", probeConfigFile(), "Config file location.")
	)

	err := fset.Parse(args)
	if err != nil {
		log.Fatalln(err)
	}

	baseConfig := loadConfigFile(*c)

	// Apply flags if they were set
	fset.Visit(func(f *flag.Flag) {
		switch f.Name {
		/*
		   case "cc":
		       baseConfig.WithCacheCredentials(*cc)
		*/
		case "hub":
			baseConfig.WithHub(*hubapi)
		case "redir":
			baseConfig.WithRedir(*redir)
		case "socks":
			baseConfig.WithSocks(*socks)
		case "doh":
			baseConfig.WithDoh(*doh)
		case "verbose":
			baseConfig.WithVerbose(*verbose)
		case "record":
			baseConfig.WithRecord(*record)
		case "insecure":
			baseConfig.WithInsecure(*insecure)
		}
	})

	// Handle positional hub argument
	if len(fset.Args()) != 0 {
		baseConfig.WithHub(fset.Args()[0])
	}

	// Set URI
	baseConfig.WithURI()

	if *version {
		printClientVersion(baseConfig)
		printHubVersion(baseConfig)
		os.Exit(0)
	}

	return baseConfig
}

// Replace Opt functions with With* methods on Config
func (c *Config) WithCacheCredentials(cc bool) *Config {
	c.CacheCredentials = cc
	return c
}

func (c *Config) WithHub(h string) *Config {
	c.Hub = h
	return c
}

func (c *Config) WithRedir(h string) *Config {
	c.Redir = h
	return c
}

func (c *Config) WithSocks(h string) *Config {
	c.Socks = h
	return c
}

func (c *Config) WithDoh(h string) *Config {
	c.Doh = h
	return c
}

func (c *Config) WithRecord(h bool) *Config {
	c.Record = h
	return c
}

func (c *Config) WithInsecure(h bool) *Config {
	c.Insecure = h
	return c
}

func (c *Config) WithURI() *Config {
	var hubapi = c.Hub
	// prepend host 127.0.0.1 to port without an explicit host :8000
	if strings.HasPrefix(hubapi, ":") {
		hubapi = "127.0.0.1" + hubapi
	}
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

func (c *Config) WithVerbose(v bool) *Config {
	c.Verbose = v
	return c
}

type clientVersion struct {
	Client *version.Version
}

type hubVersion struct {
	Hub *version.Version
}

func printClientVersion(c *Config) {
	av := &clientVersion{c.Version}
	fmt.Println(pretty.YAMLString(av))
}

func printHubVersion(c *Config) {
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

/*
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
*/
