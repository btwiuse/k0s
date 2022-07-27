package k0s

import (
	"net"
	"net/url"
	"time"
)

type Timer interface {
	Time() time.Time
}

type Namer interface {
	Name() string
}

type IDer interface {
	ID() string
}

type Tider interface {
	Namer
	Timer
	IDer
}

type Manager interface {
	Add(Tider)
	Del(string)
	Has(string) bool
	Get(string) Tider
	Keys() []string
	Values() []Tider
	Size() int
}

type Version interface {
	GetGitCommit() string
	GetGitState() string
	GetGitBranch() string
	GetGitSummary() string
	GetBuildDate() string
	GetVersion() string
	GetGoVersion() string
}

type Dialer interface {
	Dial(*url.URL) (net.Conn, error)
}
