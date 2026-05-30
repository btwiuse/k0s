//go:build js

package impl

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"k0s.io/pkg/client"
	"k0s.io/pkg/client/config"
	"k0s.io/pkg/hub/agent/info"
)

var _ client.Client = (*clientImpl)(nil)

func NewClient(c *config.Config) client.Client {
	return &clientImpl{
		config: c,
		dialer: &dialer{c},
	}
}

type clientImpl struct {
	*dialer
	config   *config.Config
	userinfo *url.Userinfo
}

func (cl *clientImpl) Config() *config.Config {
	return cl.config
}

func (cl *clientImpl) ListAgents() (agis []*info.Info, err error) {
	var (
		c  = cl.config
		ub = &url.URL{
			Scheme: c.GetScheme(),
			Host:   c.GetAddr(),
			Path:   "/api/agents/list",
		}
		ags  = []*info.Info{}
		resp *http.Response
		req  = &http.Request{
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
	resp, err = t.Do(req)
	if err != nil {
		return agis, err
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ags)
	if err != nil {
		return agis, err
	}

	return ags, err
}

func (cl *clientImpl) Run() error {
	return fmt.Errorf("interactive shell not supported in js/wasm")
}

func (cl *clientImpl) MiniRun() error {
	return fmt.Errorf("terminal not supported in js/wasm")
}

func (cl *clientImpl) RunRedir() error {
	return fmt.Errorf("redir proxy not supported in js/wasm")
}

func (cl *clientImpl) RunSocks() error {
	return fmt.Errorf("socks5 proxy not supported in js/wasm")
}

func (cl *clientImpl) RunDoh() error {
	return fmt.Errorf("DoH proxy not supported in js/wasm")
}
