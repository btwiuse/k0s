package client

import (
	"k0s.io/pkg/client/config"
	"k0s.io/pkg/hub/agent/info"
)

type Client interface {
	Config() *config.Config

	RunRedir() error
	RunSocks() error
	RunDoh() error
	Run() error
	MiniRun() error
	ListAgents() ([]*info.Info, error)
}
