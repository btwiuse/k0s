package sysinfo

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	hl "gitlab.com/mjwhitta/hilighter"
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
	var r *regexp.Regexp

	s.CPU = s.exec("sysctl", "-n", "machdep.cpu.brand_string")

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
		if (len(words) == 9) && (words[8] == path) {
			return words[2] + " / " + words[1] + " (" + words[4] + ")"
		}
	}

	return ""
}

func (s *SysInfo) kernel() string {
	s.Kernel = s.exec("sysctl", "-n", "kern.osrelease")
	return s.Kernel
}

func (s *SysInfo) operatingSystem() string {
	s.OS = s.exec("uname", "-m", "-s")
	return s.OS
}

func (s *SysInfo) ram() string {
	var e error
	var phys int
	var tmp string
	var total int
	var used int
	var user int

	tmp = s.exec("sysctl", "-n", "hw.physmem")
	if phys, e = strconv.Atoi(tmp); e != nil {
		return s.RAM
	}

	tmp = s.exec("sysctl", "-n", "hw.usermem")
	if user, e = strconv.Atoi(tmp); e != nil {
		return s.RAM
	}

	tmp = s.exec("sysctl", "-n", "hw.memsize")
	if total, e = strconv.Atoi(tmp); e != nil {
		return s.RAM
	}

	used = (phys + user) / (1024 * 1024)
	total /= 1024 * 1024

	s.RAM = strconv.Itoa(used) + " MB"
	s.RAM += " / "
	s.RAM += strconv.Itoa(total) + " MB"

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
	s.TTY = os.Getenv("GPG_TTY") // There's probably a better way
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

	return s.Uptime
}
