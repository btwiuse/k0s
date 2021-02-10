package info

import (
	"os"
	"os/user"
	"runtime"

	"k0s.io/k0s/pkg/agent"
	"k0s.io/k0s/pkg/distro"
)

type info struct {
	OS       string `json:"os"`
	Pwd      string `json:"pwd"`
	Arch     string `json:"arch"`
	Distro   string `json:"distro,omitempty"`
	Username string `json:"username"`
	Hostname string `json:"hostname"`
}

func (inf *info) GetOS() string {
	return inf.OS
}

func (inf *info) GetPwd() string {
	return inf.Pwd
}

func (inf *info) GetArch() string {
	return inf.Arch
}

func (inf *info) GetUsername() string {
	return inf.Username
}

func (inf *info) GetHostname() string {
	return inf.Hostname
}

func (inf *info) GetDistro() string {
	return inf.Distro
}

func EmptyInfo() agent.Info {
	return &info{}
}

func CollectInfo() agent.Info {
	_user, err := user.Current()
	if err != nil {
		_user = &user.User{Username: "N/A"}
	}
	var (
		pwd, _      = os.Getwd()
		username    = _user.Username
		hostname, _ = os.Hostname()
		goos        = runtime.GOOS
		goarch      = runtime.GOARCH
		distro      = distro.Vendor()
	)
	return &info{
		OS:       goos,
		Pwd:      pwd,
		Arch:     goarch,
		Distro:   distro,
		Username: username,
		Hostname: hostname,
	}
}
