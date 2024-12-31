package info

import (
	"os"
	"os/user"
	"runtime"
)

type Info struct {
	OS       string `json:"os"`
	Pwd      string `json:"pwd"`
	Arch     string `json:"arch"`
	Username string `json:"username"`
	Hostname string `json:"hostname"`
}

func CollectInfo() *Info {
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
	return &Info{
		OS:       goos,
		Pwd:      pwd,
		Arch:     goarch,
		Username: username,
		Hostname: hostname,
	}
}
