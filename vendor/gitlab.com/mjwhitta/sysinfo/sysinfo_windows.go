package sysinfo

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

type clientID struct {
	UniqueProcess uintptr
	UniqueThread  uintptr
}

type objectAttrs struct {
	Length                   uintptr
	RootDirectory            uintptr
	ObjectName               uintptr
	Attributes               uintptr
	SecurityDescriptor       uintptr
	SecurityQualityOfService uintptr
}

func (s *SysInfo) colors() string {
	// TODO colors (needs hilighter support)
	return ""
}

func (s *SysInfo) cpu() string {
	var e error
	var found bool
	var k registry.Key
	var r *regexp.Regexp

	k, found, e = registry.CreateKey(
		registry.LOCAL_MACHINE,
		filepath.Join(
			"Hardware",
			"Description",
			"System",
			"CentralProcessor",
			"0",
		),
		registry.QUERY_VALUE,
	)
	if (e != nil) || !found {
		return s.CPU
	}

	s.CPU, _, e = k.GetStringValue("ProcessorNameString")
	if e != nil {
		s.CPU = ""
		return s.CPU
	}

	r = regexp.MustCompile(`\((R|TM)\)| (@|CPU)`)
	s.CPU = r.ReplaceAllString(s.CPU, "")

	r = regexp.MustCompile(`\s+`)
	s.CPU = r.ReplaceAllString(s.CPU, " ")

	return s.CPU
}

func (s *SysInfo) filesystems() []string {
	var out = []string{}

	s.RootFS = s.fsUsage("C:")

	if s.RootFS != "" {
		out = append(out, s.RootFS)
	}

	return out
}

func (s *SysInfo) fsUsage(path string) string {
	var avail int
	var e error
	var tmp []string
	var total int
	var used int

	tmp = strings.Split(
		s.exec(
			"powershell",
			"-c",
			strings.Join(
				[]string{
					"gcim Win32_LogicalDisk -Filter \"name='",
					path,
					"'\" | select Freespace,Size",
				},
				"",
			),
		),
		"\n",
	)
	if len(tmp) == 3 {
		tmp = strings.Fields(tmp[2])

		if len(tmp) == 2 {
			if avail, e = strconv.Atoi(tmp[0]); e != nil {
				return s.RAM
			}

			if total, e = strconv.Atoi(tmp[1]); e != nil {
				return s.RAM
			}

			avail /= 1024 * 1024 * 1024
			total /= 1024 * 1024 * 1024
			used = total - avail

			tmp = []string{
				strconv.Itoa(used),
				strconv.Itoa(total),
				strconv.Itoa(100 * used / total),
			}

			return tmp[0] + "G / " + tmp[1] + "G (" + tmp[2] + "%)"
		}
	}

	return ""
}

func (s *SysInfo) kernel() string {
	var e error
	var found bool
	var k registry.Key
	var minor uint64
	var tmp string

	k, found, e = registry.CreateKey(
		registry.LOCAL_MACHINE,
		filepath.Join(
			"Software",
			"Microsoft",
			"Windows NT",
			"CurrentVersion",
		),
		registry.QUERY_VALUE,
	)
	if (e != nil) || !found {
		return s.Kernel
	}

	if s.Kernel, _, e = k.GetStringValue("DisplayVersion"); e != nil {
		s.Kernel = ""
		return s.Kernel
	}

	if tmp, _, e = k.GetStringValue("CurrentBuild"); e != nil {
		return s.Kernel
	}

	s.Kernel += " (OS Build " + tmp

	if minor, _, e = k.GetIntegerValue("UBR"); e == nil {
		s.Kernel += "." + strconv.Itoa(int(minor))
	}

	s.Kernel += ")"
	return s.Kernel
}

func (s *SysInfo) operatingSystem() string {
	var e error
	var found bool
	var k registry.Key

	k, found, e = registry.CreateKey(
		registry.LOCAL_MACHINE,
		filepath.Join(
			"Software",
			"Microsoft",
			"Windows NT",
			"CurrentVersion",
		),
		registry.QUERY_VALUE,
	)
	if (e != nil) || !found {
		s.OS = "Windows"
		return s.OS
	}

	if s.OS, _, e = k.GetStringValue("ProductName"); e != nil {
		s.OS = "Windows"
	}

	return s.OS
}

func (s *SysInfo) ram() string {
	var avail int
	var e error
	var tmp string
	var total int
	var used int

	tmp = s.exec(
		"powershell",
		"-c",
		"(Get-Counter '\\Memory\\Available Bytes').CounterSamples.CookedValue",
	)
	if avail, e = strconv.Atoi(tmp); e != nil {
		return s.RAM
	}

	tmp = s.exec(
		"powershell",
		"-c",
		"(gcim Win32_PhysicalMemory | measure -Property capacity -Sum).Sum",
	)
	if total, e = strconv.Atoi(tmp); e != nil {
		return s.RAM
	}

	avail /= 1024 * 1024
	total /= 1024 * 1024
	used = total - avail

	s.RAM = strconv.Itoa(used) + " MB"
	s.RAM += " / "
	s.RAM += strconv.Itoa(total) + " MB"

	return s.RAM
}

func (s *SysInfo) shell() string {
	var err uintptr
	var lpwstrFilename []uint16 = make([]uint16, 64)
	var n uintptr
	var ntdll *windows.LazyDLL = windows.NewLazySystemDLL("ntdll")
	var pHndl windows.Handle
	var ppid int = os.Getppid()
	var psapi *windows.LazyDLL = windows.NewLazySystemDLL("psapi")

	err, _, _ = ntdll.NewProc("NtOpenProcess").Call(
		uintptr(unsafe.Pointer(&pHndl)),
		uintptr(0x0400), // PROCESS_QUERY_INFORMATION
		uintptr(unsafe.Pointer(&objectAttrs{0, 0, 0, 0, 0, 0})),
		uintptr(unsafe.Pointer(&clientID{uintptr(ppid), 0})),
	)
	if (err != 0) || (pHndl == 0) {
		return s.Shell
	}

	n, _, _ = psapi.NewProc("GetModuleFileNameExW").Call(
		uintptr(pHndl),
		0,
		uintptr(unsafe.Pointer(&lpwstrFilename[0])),
		64,
	)
	if n == 0 {
		return s.Shell
	}

	s.Shell = filepath.Base(windows.UTF16ToString(lpwstrFilename[:n]))
	return s.Shell
}

func (s *SysInfo) tty() string {
	return s.TTY
}

func (s *SysInfo) uptime() string {
	var out string = s.exec(
		"powershell",
		"-c",
		"(date) - (gcim Win32_OperatingSystem).LastBootUpTime",
	)
	var tmp []string

	for _, line := range strings.Split(out, "\n") {
		if strings.HasPrefix(line, "Days") {
			tmp = strings.Fields(line)

			if len(tmp) >= 3 {
				if tmp[len(tmp)-1] != "0" {
					s.Uptime = tmp[len(tmp)-1] + " days"
				}
			}
		} else if strings.HasPrefix(line, "Hours") {
			tmp = strings.Fields(line)

			if len(tmp) >= 3 {
				if tmp[len(tmp)-1] != "0" {
					if s.Uptime != "" {
						s.Uptime += ", "
					}

					s.Uptime += tmp[len(tmp)-1] + " hours"
				}
			}
		} else if strings.HasPrefix(line, "Minutes") {
			tmp = strings.Fields(line)

			if len(tmp) >= 3 {
				if tmp[len(tmp)-1] != "0" {
					if s.Uptime != "" {
						s.Uptime += ", "
					}

					s.Uptime += tmp[len(tmp)-1] + " mins"
				}
			}
		}
	}

	return s.Uptime
}
