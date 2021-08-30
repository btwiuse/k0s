//go:build !darwin && !windows
// +build !darwin,!windows

package sysinfo

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	hl "gitlab.com/mjwhitta/hilighter"
	"gitlab.com/mjwhitta/pathname"
)

func (s *SysInfo) colors() string {
	s.Colors = strings.Join(
		[]string{
			hl.Hilights([]string{"light_black", "on_black"}, "▄▄▄"),
			hl.Hilights([]string{"light_red", "on_red"}, "▄▄▄"),
			hl.Hilights([]string{"light_green", "on_green"}, "▄▄▄"),
			hl.Hilights([]string{"light_yellow", "on_yellow"}, "▄▄▄"),
			hl.Hilights([]string{"light_blue", "on_blue"}, "▄▄▄"),
			hl.Hilights(
				[]string{"light_magenta", "on_magenta"},
				"▄▄▄",
			),
			hl.Hilights([]string{"light_cyan", "on_cyan"}, "▄▄▄"),
			hl.Hilights([]string{"light_white", "on_white"}, "▄▄▄"),
		},
		"",
	)
	return s.Colors
}

func (s *SysInfo) cpu() string {
	var e error
	var info []byte
	var r *regexp.Regexp

	if info, e = ioutil.ReadFile("/proc/cpuinfo"); e != nil {
		s.CPU = ""
		return s.CPU
	}

	r = regexp.MustCompile(`(cpu model|model name)\s+:\s+(.+)`)
	for _, match := range r.FindAllStringSubmatch(string(info), -1) {
		s.CPU = match[2]
		break
	}

	r = regexp.MustCompile(`\((R|TM)\)| (@|CPU)`)
	s.CPU = r.ReplaceAllString(s.CPU, "")

	r = regexp.MustCompile(`\s+`)
	s.CPU = r.ReplaceAllString(s.CPU, " ")

	return s.CPU
}

func (s *SysInfo) filesystems() []string {
	var out = []string{}

	s.RootFS = s.fsUsage("/")
	s.HomeFS = s.fsUsage("/home")

	if s.RootFS != "" {
		out = append(out, s.RootFS)
	}

	if (s.HomeFS != "") && (s.HomeFS != s.RootFS) {
		out = append(out, s.HomeFS)
	} else {
		s.HomeFS = ""
	}

	return out
}

func (s *SysInfo) fsUsage(path string) string {
	var usage string
	var words []string

	usage = s.exec("df", "-h", path)

	for _, line := range strings.Split(usage, "\n") {
		words = strings.Fields(line)
		if (len(words) == 6) && (words[5] == path) {
			return words[2] + " / " + words[1] + " (" + words[4] + ")"
		}
	}

	return ""
}

func (s *SysInfo) kernel() string {
	var e error
	var kernel []byte

	kernel, e = ioutil.ReadFile("/proc/sys/kernel/osrelease")
	if e == nil {
		s.Kernel = strings.TrimSpace(string(kernel))
	}

	return s.Kernel
}

func (s *SysInfo) operatingSystem() string {
	var e error
	var matches [][]string
	var r *regexp.Regexp
	var release []byte

	if pathname.DoesExist("/etc/os-release") {
		if release, e = ioutil.ReadFile("/etc/os-release"); e != nil {
			s.OS = ""
			return s.OS
		}

		r = regexp.MustCompile(`PRETTY_NAME="(.+)"`)
		matches = r.FindAllStringSubmatch(string(release), -1)
		for _, match := range matches {
			s.OS = match[1] + " " + s.exec("uname", "-m")
			break
		}
	} else {
		s.OS = s.exec("uname", "-m", "-s")
	}

	return s.OS
}

func (s *SysInfo) ram() string {
	var matches [][]string
	var r *regexp.Regexp
	var total int
	var used int

	r = regexp.MustCompile(`Mem:\s+(\d+)\s+(\d+)`)
	matches = r.FindAllStringSubmatch(s.exec("free"), -1)
	for _, match := range matches {
		total, _ = strconv.Atoi(match[1])
		used, _ = strconv.Atoi(match[2])

		total /= 1024
		used /= 1024

		s.RAM = strconv.Itoa(used) + " MB"
		s.RAM += " / "
		s.RAM += strconv.Itoa(total) + " MB"

		break
	}

	return s.RAM
}

func (s *SysInfo) shell() string {
	var exists bool
	var sh string

	if sh, exists = os.LookupEnv("SHELL"); exists {
		s.Shell = strings.TrimSpace(sh)
	}

	return s.Shell
}

func (s *SysInfo) tty() string {
	var e error
	if s.TTY, e = os.Readlink("/proc/self/fd/0"); e != nil {
		s.TTY = ""
	}
	return s.TTY
}

func (s *SysInfo) uptime() string {
	var r *regexp.Regexp

	s.Uptime = s.exec("uptime")

	// Fail fast
	if s.Uptime == "" {
		return s.Uptime
	}

	// Strip extra whitespace
	r = regexp.MustCompile(`\s+`)
	s.Uptime = r.ReplaceAllString(s.Uptime, " ")

	// Strip leading and trailing data
	r = regexp.MustCompile(`^.*up\s+|,\s+\d+\s+user.+$`)
	s.Uptime = r.ReplaceAllString(s.Uptime, "")

	// Convert hours:mins to match days
	r = regexp.MustCompile(`0?(\d+):0?(\d+)`)
	s.Uptime = r.ReplaceAllString(s.Uptime, "$1 hour, $2 min")

	// Remove 0 hours and mins
	r = regexp.MustCompile(`(^|,\s+)0\s+(hour|min)`)
	s.Uptime = r.ReplaceAllString(s.Uptime, "")

	// Make plural if needed
	r = regexp.MustCompile(`((\d\d+|[2-9])\s+(hour|min))`)
	s.Uptime = r.ReplaceAllString(s.Uptime, "${1}s")

	// Remove leading and trailing commas or whitespace
	r = regexp.MustCompile(`^(,?\s*)+|(,?\s*)+$`)
	s.Uptime = r.ReplaceAllString(s.Uptime, "")

	if s.Uptime == "" {
		return "0 mins"
	}

	return s.Uptime
}
