package info

import (
	"os"
	"os/user"
	"runtime"

	"k0s.io/pkg/agent"
)

type info struct {
	OS       string `json:"os"`
	Pwd      string `json:"pwd"`
	Arch     string `json:"arch"`
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
	)
	return &info{
		OS:       goos,
		Pwd:      pwd,
		Arch:     goarch,
		Username: username,
		Hostname: hostname,
	}
}
