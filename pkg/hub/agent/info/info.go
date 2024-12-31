package info

import (
	"encoding/json"
	"time"

	"github.com/btwiuse/version"
)

type Meta struct {
	OS       string `json:"os"`
	Pwd      string `json:"pwd"`
	Arch     string `json:"arch"`
	Username string `json:"username"`
	Hostname string `json:"hostname"`
}

// only Unmarshed from agent header
// never Marshaled
type privateInfo struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Tags     []string          `json:"tags"`
	Htpasswd map[string]string `json:"htpasswd,omitempty"`

	Meta             `json:"meta"`
	*version.Version `json:"version"`
}

// Metadata, for flatten json output
// Marshed by hub /api/agents/list
// Unmarshaled by client
type Info struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Tags      []string `json:"tags"`
	Auth      *bool    `json:"auth,omitempty"`
	Connected int64    `json:"connected"`
	IP        string   `json:"ip"`

	OS         string `json:"os"`
	Pwd        string `json:"pwd"`
	Arch       string `json:"arch"`
	Username   string `json:"username"`
	Hostname   string `json:"hostname"`
	Version    string `json:"version"`
	GitSummary string `json:"git_summary"`

	*privateInfo `json:"-"`
}

func Decode(data []byte) (*Info, error) {
	pi := &privateInfo{}
	err := json.Unmarshal(data, pi)
	if err != nil {
		return nil, err
	}

	i := &Info{privateInfo: pi}
	i.populatePublicInfo()
	return i, err
}

func (info *Info) populatePublicInfo() {
	pi := info.privateInfo

	info.ID = pi.ID
	info.Name = pi.Name
	info.Tags = pi.Tags
	info.Connected = time.Now().Unix()
	info.Auth = new(bool)

	info.OS = pi.OS
	info.Pwd = pi.Pwd
	info.Arch = pi.Arch
	info.Username = pi.Username
	info.Hostname = pi.Hostname
	info.GitSummary = pi.Version.GitVersion

	if len(pi.Htpasswd) != 0 {
		*info.Auth = true
	}
}

func (info *Info) SetIP(ip string) {
	info.IP = ip
}
