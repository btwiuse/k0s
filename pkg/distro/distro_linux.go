//go:build (linux || freebsd || openbsd || netbsd) && !windows && !darwin

package distro

import (
	"github.com/mjwhitta/sysinfo"
)

var Info = sysinfo.New()

func Vendor() string {
	return "linux"
}

func Name() string {
	return Info.OS
}
