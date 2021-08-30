// +build !windows
// +build !linux
// +build darwin

package distro

import (
	"gitlab.com/mjwhitta/sysinfo"
)

var Info = sysinfo.New()

func Vendor() string {
	return "darwin"
}

func Name() string {
	return Info.OS
}
