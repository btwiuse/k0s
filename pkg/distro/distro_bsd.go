//go:build (freebsd || openbsd || netbsd) && !windows && !darwin && !linux

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
