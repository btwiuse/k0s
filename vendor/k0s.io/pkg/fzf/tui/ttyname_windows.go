//go:build windows
// +build windows

package tui

func ttyname() string {
	return ""
}
