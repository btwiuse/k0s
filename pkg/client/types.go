package client

import (
	"net"

	"k0s.io/conntroll/pkg"
)

type Client interface {
	Config

	RunRedir() error
	RunSocks() error
	RunSocks5ToHTTP() error
	Run() error
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
	Config

	GetID() string
	GetName() string
	GetTags() []string

	GetCmd() []string
	GetReadOnly() bool
	GetVerbose() bool

	String() string
}

type Dialer interface {
	Dial(string) (net.Conn, error)
}
