package client

import (
	"net"

	"k0s.io/conntroll/pkg"
)

type Client interface {
	Config

	RunRedir() error
	RunSocks() error
	Run() error
}

type Config interface {
	GetHost() string
	GetPort() string
	GetAddr() string
	GetScheme() string

	GetSocks() string
	GetRedir() string

	GetVersion() pkg.Version
	GetInsecure() bool
}

type Info interface {
	GetOS() string
	GetPwd() string
	GetArch() string
	GetDistro() string
	GetHostname() string
	GetUsername() string
}

type Agent interface {
	Info

	GetID() string
	GetName() string
	GetTags() []string

	GetHost() string
	GetPort() string
	GetAddr() string
	GetScheme() string

	GetCmd() []string
	GetReadOnly() bool
	GetVerbose() bool
	GetInsecure() bool

	String() string

	GetVersion() pkg.Version
}

type Dialer interface {
	Dial(string) (net.Conn, error)
}
