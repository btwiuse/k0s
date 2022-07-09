//go:build windows && !linux && !darwin
// +build windows,!linux,!darwin

package distro

func Vendor() string {
	return "windows"
}

func Name() string {
	return "Microsoft Windows"
}
