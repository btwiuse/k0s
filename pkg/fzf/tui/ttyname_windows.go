//go:build windows

package tui

func ttyname() string {
	return ""
}
