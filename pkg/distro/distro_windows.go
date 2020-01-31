// +build windows
// +build !linux
// +build !darwin

package distro

func Vendor() string {
	return "windows"
}

func Name() string {
	return "Microsoft Windows"
}
