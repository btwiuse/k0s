package client

import (
	"net"

	"k0s.io/conntroll/pkg"
	"k0s.io/conntroll/pkg/hub"
)

type Client interface {
	Config

	RunRedir() error
	RunSocks() error
	RunSocks5ToHTTP() error
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
	GetSocks5ToHTTP() string
	GetCacheCredentials() bool
	GetCredentials() map[string]KeyStore
	GetConfigLocation() string

	GetVersion() pkg.Version
	GetInsecure() bool
}

type Dialer interface {
	Dial(string) (net.Conn, error)
}

type KeyStore map[string]string
