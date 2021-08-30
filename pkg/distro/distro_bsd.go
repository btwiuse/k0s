// +build freebsd openbsd netbsd
// +build !windows
// +build !darwin
// +build !linux

package distro

import (
	"gitlab.com/mjwhitta/sysinfo"
)

var Info = sysinfo.New()

func Vendor() string {
	return "bsd"
}

func Name() string {
	return Info.OS
}
