package info

import (
	"encoding/json"

	"github.com/btwiuse/conntroll/pkg/hub"
)

var _ hub.Info = (*Info)(nil)

type Meta struct {
	OS       string `json:"os"`
	Pwd      string `json:"pwd"`
	Arch     string `json:"arch"`
	Distro   string `json:"distro,omitempty"`
	Username string `json:"username"`
	Hostname string `json:"hostname"`
}

type Info struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
	Auth string `json:"auth,omitempty"`

	Meta `json:"meta"`
}

func Decode(data []byte) (*Info, error) {
	v := &Info{
		// Meta: &Meta{},
	}
	err := json.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (info *Info) GetOS() string{
	return info.OS
}

func (info *Info) GetPwd() string{
	return info.Pwd
}

func (info *Info) GetArch() string{
	return info.Arch
}

func (info *Info) GetDistro() string{
	return info.Distro
}

func (info *Info) GetUsername() string{
	return info.Username
}

func (info *Info) GetHostname() string{
	return info.Hostname
}

func (info *Info) GetID() string{
	return info.ID
}

func (info *Info) GetName() string{
	return info.Name
}

func (info *Info) GetTags() []string{
	return info.Tags
}

func (info *Info) GetAuth() string{
	return info.Auth
}
