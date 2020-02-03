// +build bin

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mbndr/figlet4go"
)

func figlet(str string) {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render(str)
	fmt.Print(renderStr)
}

func (c Combo) String() string {
	return fmt.Sprintf("%s/%s", c.OS, c.ARCH)
}

type Combo struct {
	OS   string
	ARCH string
}

var (
	Basename     = "k0s"
	Path         = "bin"
	Delimeter    = "/"
	DefaultCombo = Combo{runtime.GOOS, runtime.GOARCH}
	GlobalEnv    = []string{
		"CGO_ENABLED=0",
	}
)

func (c Combo) ReleaseName() string {
	exe := Basename
	if c.OS == "windows" {
		exe = fmt.Sprintf("%s.exe", exe)
	}
	return strings.Join([]string{c.OS, c.ARCH, exe}, Delimeter)
}

func (c Combo) Env() []string {
	envs := append(GlobalEnv, fmt.Sprintf("GOOS=%s", c.OS))
	if strings.HasPrefix(c.ARCH, "armv") {
		v := strings.TrimPrefix(c.ARCH, "armv")
		envs = append(envs, fmt.Sprintf("GOARCH=%s", "arm"), fmt.Sprintf("GOARM=%s", v))
	} else {
		envs = append(envs, fmt.Sprintf("GOARCH=%s", c.ARCH))
	}
	switch c {
	// currently I haven't figured out how to properly setup the toolchain for darwin/arm*
	case
		Combo{OS: "darwin", ARCH: "arm64"},
		Combo{OS: "darwin", ARCH: "armv6"},
		Combo{OS: "darwin", ARCH: "armv7"}:
		envs = append(envs, "CGO_ENABLED=1")
	case Combo{OS: "android", ARCH: "arm64"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=aarch64-linux-android29-clang",
			"CXX=aarch64-linux-android29-clang++",
		)
	case Combo{OS: "android", ARCH: "armv7"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=armv7a-linux-androideabi29-clang",
			"CXX=armv7a-linux-androideabi29-clang++",
		)
	case Combo{OS: "android", ARCH: "armv6"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=armv7a-linux-androideabi29-clang",
			"CXX=armv7a-linux-androideabi29-clang++",
		)
	case Combo{OS: "android", ARCH: "amd64"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=x86_64-linux-android29-clang",
			"CXX=x86_64-linux-android29-clang++",
		)
	case Combo{OS: "android", ARCH: "386"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=i686-linux-android29-clang",
			"CXX=i686-linux-android29-clang++",
		)
	case Combo{OS: "windows", ARCH: "amd64"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CXX=x86_64-w64-mingw32-g++",
			"CC=x86_64-w64-mingw32-gcc",
		)
	case Combo{OS: "windows", ARCH: "386"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CXX=i686-w64-mingw32-g++",
			"CC=i686-w64-mingw32-gcc",
		)
	// windows/arm doesn't work yet
	case Combo{OS: "windows", ARCH: "arm"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CXX=clang++",
			"CC=clang",
		)
	case Combo{OS: "linux", ARCH: "armv7"}:
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=arm-linux-gnueabihf-gcc",
			"CXX=arm-linux-gnueabihf-g++",
		)
	case Combo{OS: "android", ARCH: "armv6"}:
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=arm-linux-gnueabihf-gcc",
			"CXX=arm-linux-gnueabihf-g++",
		)
	}
	return envs
}

func parseCombo(osarch string) Combo {
	parts := strings.Split(osarch, "/")
	if len(parts) != 2 {
		log.Fatalln("error parsing combo", osarch)
	}
	return Combo{
		OS:   parts[0],
		ARCH: parts[1],
	}
}

func main() {
	var (
		stripFlag bool
		upxFlag   bool
		tags      string
		ldflags   string
	)

	flag.StringVar(&Path, "d", Path, "output directory")
	flag.StringVar(&tags, "tags", "", "build tags")
	flag.StringVar(&ldflags, "ldflags", "", "ldflags")
	flag.BoolVar(&stripFlag, "strip", false, "strip binary")
	flag.BoolVar(&upxFlag, "upx", false, "compress binary with upx")
	flag.Parse()

	combos := []Combo{}
	for _, arg := range flag.Args() {
		combos = append(combos, parseCombo(arg))
	}

	if len(combos) == 0 {
		combos = append(combos, DefaultCombo)
	}

	for _, c := range combos {
		figlet(c.String())
		var (
			buildArgs = []string{"build",
				"-o", filepath.Join(Path, c.ReleaseName()),
				"-mod=vendor",
				"-trimpath",
				"-ldflags", ldflags,
				"-tags", tags,
				"-v",
				".",
			}
			build = exec.Command("go", buildArgs...)
			strip = exec.Command(
				"strip",
				filepath.Join(Path, c.ReleaseName()),
			)
			upx = exec.Command(
				"upx",
				filepath.Join(Path, c.ReleaseName()),
			)
		)

		build.Env = append(os.Environ(), c.Env()...)
		build.Stdout = os.Stdout
		build.Stderr = os.Stderr
		log.Println("Env:", c.Env())
		for i, arg := range buildArgs {
			log.Println(fmt.Sprintf("%d %q", i, arg))
		}

		if err := build.Run(); err != nil {
			log.Fatalln(err)
		}

		// strip fails on arm64 binary, here we simply ignore it
		if stripFlag && c.OS == "linux" {
			switch c.ARCH {
			case "arm64", "mips", "mipsle", "mips64", "mips64le", "s390x":
				break
			default:
				strip.Run()
			}
		}

		if upxFlag && c.ARCH != "arm64" {
			upx.Stdout = os.Stdout
			upx.Stderr = os.Stderr
			if err := upx.Run(); err != nil {
				log.Fatalln(err)
			}
		}

		if c == DefaultCombo {
			for _, bin := range []string{"k0s"} {
				src := filepath.Join(c.ReleaseName())
				dst := filepath.Join(Path, bin)
				ln(src, dst)
			}
		}
	}
}

func ln(from, to string) {
	lnf := exec.Command("ln", "-s", "-f", "-v", from, to)
	lnf.Stdout = os.Stdout
	lnf.Stderr = os.Stderr
	if err := lnf.Run(); err != nil {
		log.Fatalln(err)
	}
}
