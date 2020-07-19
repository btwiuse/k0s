package client

import (
	"net"
	"net/url"

	"k0s.io/k0s/pkg"
	"k0s.io/k0s/pkg/hub"
)

type Client interface {
	Config

	RunRedir() error
	RunSocks() error
	Run() error
	ListAgents() ([]hub.AgentInfo, error)
}

type Config interface {
	GetHost() string
	GetPort() string
	GetAddr() string
	GetScheme() string
	GetSchemeWS() string

	GetRedir() string
	GetSocks() string
	GetCacheCredentials() bool
	GetCredentials() map[string]KeyStore
	GetConfigLocation() string

	GetVersion() pkg.Version
	GetInsecure() bool
	GetRecord() bool
}

type Dialer interface {
	Dial(string, *url.Userinfo) (net.Conn, error)
}

type KeyStore map[string]string
