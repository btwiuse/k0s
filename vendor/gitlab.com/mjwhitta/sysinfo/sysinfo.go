package sysinfo

import (
	"encoding/json"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"

	hl "gitlab.com/mjwhitta/hilighter"
	"gitlab.com/mjwhitta/where"
)

// SysInfo is a struct containing relevant system information.
type SysInfo struct {
	Colors      string `json:"-"`
	CPU         string `json:"CPU,omitempty"`
	dataColors  []string
	fieldColors []string
	Height      int    `json:"-"`
	HomeFS      string `json:"HomeFS,omitempty"`
	Host        string `json:"Host,omitempty"`
	ipMutex     *sync.Mutex
	ips         map[string][]string
	IPv4        string `json:"IPv4,omitempty"`
	IPv6        string `json:"IPv6,omitempty"`
	Kernel      string `json:"Kernel,omitempty"`
	order       []string
	OS          string `json:"OS,omitempty"`
	RAM         string `json:"RAM,omitempty"`
	RootFS      string `json:"RootFS,omitempty"`
	Shell       string `json:"Shell,omitempty"`
	TTY         string `json:"TTY,omitempty"`
	Uptime      string `json:"Uptime,omitempty"`
	Width       int    `json:"-"`
}

// New will return a SysInfo pointer. A list of fields can be
// supplied if all info is not wanted.
func New(fields ...string) *SysInfo {
	var s = &SysInfo{ipMutex: &sync.Mutex{}}

	s.order = fields
	if len(fields) == 0 {
		s.order = []string{
			"host",
			"os",
			"kernel",
			"uptime",
			"ipv4",
			"ipv6",
			"shell",
			"tty",
			"cpu",
			"ram",
			"fs",
			"blank",
			"colors",
		}
	}

	s.Collect()

	return s
}

func (s *SysInfo) calcSize() {
	s.Height = 0
	s.Width = 0

	for _, line := range strings.Split(hl.Plain(s.String()), "\n") {
		s.Height++
		if len([]rune(line)) > s.Width {
			s.Width = len([]rune(line))
		}
	}
}

// Clear will remove all system info.
func (s *SysInfo) Clear() {
	s.Colors = ""
	s.CPU = ""
	s.HomeFS = ""
	s.Host = ""
	s.ips = nil
	s.IPv4 = ""
	s.IPv6 = ""
	s.Kernel = ""
	s.OS = ""
	s.RAM = ""
	s.RootFS = ""
	s.Shell = ""
	s.TTY = ""
	s.Uptime = ""
	s.calcSize()
}

// Collect will get requested system info.
func (s *SysInfo) Collect() {
	var newOrder []string
	var wg = sync.WaitGroup{}

	for _, field := range s.order {
		switch strings.ToLower(field) {
		case "blank":
			newOrder = append(newOrder, "Blank")
		case "colors":
			newOrder = append(newOrder, "Colors")

			wg.Add(1)
			go func() {
				s.colors()
				wg.Done()
			}()
		case "cpu":
			newOrder = append(newOrder, "CPU")

			wg.Add(1)
			go func() {
				s.cpu()
				wg.Done()
			}()
		case "fs":
			newOrder = append(newOrder, "FS")

			wg.Add(1)
			go func() {
				s.filesystems()
				wg.Done()
			}()
		case "host":
			newOrder = append(newOrder, "Host")

			wg.Add(1)
			go func() {
				s.hostname()
				wg.Done()
			}()
		case "ipv4":
			newOrder = append(newOrder, "IPv4")

			wg.Add(1)
			go func() {
				s.ipv4()
				wg.Done()
			}()
		case "ipv6":
			newOrder = append(newOrder, "IPv6")

			wg.Add(1)
			go func() {
				s.ipv6()
				wg.Done()
			}()
		case "kernel":
			newOrder = append(newOrder, "Kernel")

			wg.Add(1)
			go func() {
				s.kernel()
				wg.Done()
			}()
		case "os":
			newOrder = append(newOrder, "OS")

			wg.Add(1)
			go func() {
				s.operatingSystem()
				wg.Done()
			}()
		case "ram":
			newOrder = append(newOrder, "RAM")

			wg.Add(1)
			go func() {
				s.ram()
				wg.Done()
			}()
		case "shell":
			newOrder = append(newOrder, "Shell")

			wg.Add(1)
			go func() {
				s.shell()
				wg.Done()
			}()
		case "tty":
			newOrder = append(newOrder, "TTY")

			wg.Add(1)
			go func() {
				s.tty()
				wg.Done()
			}()
		case "uptime":
			newOrder = append(newOrder, "Uptime")

			wg.Add(1)
			go func() {
				s.uptime()
				wg.Done()
			}()
		}
	}

	s.order = newOrder
	wg.Wait()
	s.calcSize()
}

func (s *SysInfo) exec(cmd string, cli ...string) string {
	var e error
	var o []byte

	if (cmd == "") || (where.Is(cmd) == "") {
		return ""
	}

	if o, e = exec.Command(cmd, cli...).Output(); e != nil {
		return ""
	}

	return strings.TrimSpace(string(o))
}

func (s *SysInfo) format(k string, v string, max int) string {
	var filler string
	var line string

	for i := 0; i < max-len(k); i++ {
		filler += " "
	}

	line = " " + hl.Hilights(s.fieldColors, filler+k+":")
	line += " "
	line += hl.Hilights(s.dataColors, v)

	return line
}

func (s *SysInfo) getIPs() map[string][]string {
	var addrs []net.Addr
	var e error
	var ifaces []net.Interface
	var ip net.IP
	var tmp string

	s.ipMutex.Lock()
	defer s.ipMutex.Unlock()

	if (s.ips != nil) && (len(s.ips) > 0) {
		return s.ips
	}

	s.ips = map[string][]string{}

	if ifaces, e = net.Interfaces(); e != nil {
		return s.ips
	}

	for _, iface := range ifaces {
		if addrs, e = iface.Addrs(); e != nil {
			continue
		}

		for _, addr := range addrs {
			tmp = addr.String()
			tmp = tmp[0:strings.Index(tmp, "/")]

			if ip = net.ParseIP(tmp); (ip == nil) || ip.IsLoopback() {
				continue
			}

			if ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
				continue
			}

			if strings.HasPrefix(iface.Name, "docker") {
				continue
			}

			s.ips[iface.Name] = append(
				s.ips[iface.Name],
				addr.String(),
			)
		}
	}

	return s.ips
}

func (s *SysInfo) hostname() string {
	var e error

	if s.Host, e = os.Hostname(); e != nil {
		s.Host = ""
	}

	return s.Host
}

func (s *SysInfo) ipv4() string {
	var ip net.IP
	var tmp string

	for _, vs := range s.getIPs() {
		for _, v := range vs {
			tmp = v[0:strings.Index(v, "/")]

			if ip = net.ParseIP(tmp); ip == nil {
				continue
			}

			if ip.To4() != nil {
				s.IPv4 = v
				return s.IPv4
			}
		}
	}

	return s.IPv4
}

func (s *SysInfo) ipv6() string {
	var ip net.IP
	var tmp string

	for _, vs := range s.getIPs() {
		for _, v := range vs {
			tmp = v[0:strings.Index(v, "/")]

			if ip = net.ParseIP(tmp); ip == nil {
				continue
			}

			if ip.To4() == nil {
				s.IPv6 = v
				return s.IPv6
			}
		}
	}

	return s.IPv6
}

// SetDataColors will set the color values for the field data. See
// gitlab.com/mjwhitta/hilighter for valid colors.
func (s *SysInfo) SetDataColors(colors ...string) {
	s.dataColors = colors
}

// SetFieldColors will set the color values for the field names. See
// gitlab.com/mjwhitta/hilighter for valid colors.
func (s *SysInfo) SetFieldColors(colors ...string) {
	s.fieldColors = colors
}

// String will convert the SysInfo struct to a printable string.
func (s *SysInfo) String() string {
	var data = map[string]string{}
	var hasKey bool
	var max int
	var out []string
	var tmp []byte

	tmp, _ = json.Marshal(s)
	json.Unmarshal(tmp, &data)

	for k := range data {
		if len(k) > max {
			max = len(k)
		}
	}

	for _, field := range s.order {
		switch field {
		case "Blank":
			out = append(out, "")
		case "Colors":
			if s.Colors != "" {
				out = append(out, " "+s.Colors)
			}
		case "FS":
			field = "RootFS"
			if _, hasKey = data[field]; hasKey {
				out = append(out, s.format(field, data[field], max))
			}

			field = "HomeFS"
			if _, hasKey = data[field]; hasKey {
				out = append(out, s.format(field, data[field], max))
			}
		default:
			if _, hasKey = data[field]; hasKey {
				out = append(out, s.format(field, data[field], max))
			}
		}
	}

	return strings.Join(out, "\n")
}
