package client

import (
	"k0s.io/pkg/hub"
	"github.com/btwiuse/version"
)

type Client interface {
	Config

	RunRedir() error
	RunSocks() error
	RunDoh() error
	Run() error
	MiniRun() error
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
	GetDoh() string
	GetCacheCredentials() bool
	GetCredentials() map[string]KeyStore
	GetConfigLocation() string

	GetVersion() *version.Version
	GetInsecure() bool
	GetRecord() bool
}

type KeyStore map[string]string
