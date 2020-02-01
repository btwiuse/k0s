// +build linux darwin
// +build !windows

package distro

import (
	"github.com/pupapaik/sysinfo"
)

func Vendor() string {
	si := &sysinfo.SysInfo{}
	si.GetSysInfo()
	return si.OS.Vendor
}

func Name() string {
	si := &sysinfo.SysInfo{}
	si.GetSysInfo()
	return si.OS.Name
}
