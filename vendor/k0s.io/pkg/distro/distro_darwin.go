//go:build !windows && !linux && darwin
// +build !windows,!linux,darwin

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
