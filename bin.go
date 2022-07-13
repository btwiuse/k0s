//go:build bin
// +build bin

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"k0s.io/pkg/fonts"
)

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
	//                  _           _     _
	//   __ _ _ __   __| |_ __ ___ (_) __| |
	//  / _` | '_ \ / _` | '__/ _ \| |/ _` |
	// | (_| | | | | (_| | | | (_) | | (_| |
	//  \__,_|_| |_|\__,_|_|  \___/|_|\__,_|
	//
	case Combo{OS: "android", ARCH: "armv6"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=armv7a-linux-androideabi30-clang",
			"CXX=armv7a-linux-androideabi30-clang++",
		)
	case Combo{OS: "android", ARCH: "armv7"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=armv7a-linux-androideabi30-clang",
			"CXX=armv7a-linux-androideabi30-clang++",
		)
	case Combo{OS: "android", ARCH: "arm64"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=aarch64-linux-android30-clang",
			"CXX=aarch64-linux-android30-clang++",
		)
	case Combo{OS: "android", ARCH: "386"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=i686-linux-android30-clang",
			"CXX=i686-linux-android30-clang++",
		)
	case Combo{OS: "android", ARCH: "amd64"}:
		envs = append(envs, "CGO_ENABLED=1")
		if c == DefaultCombo {
			break
		}
		envs = append(envs,
			"CC=x86_64-linux-android30-clang",
			"CXX=x86_64-linux-android30-clang++",
		)

	//           _           _
	// __      _(_)_ __   __| | _____      _____
	// \ \ /\ / / | '_ \ / _` |/ _ \ \ /\ / / __|
	//  \ V  V /| | | | | (_| | (_) \ V  V /\__ \
	//   \_/\_/ |_|_| |_|\__,_|\___/ \_/\_/ |___/
	//

	//  _ _
	// | (_)_ __  _   ___  __
	// | | | '_ \| | | \ \/ /
	// | | | | | | |_| |>  <
	// |_|_|_| |_|\__,_/_/\_\
	//
	case Combo{OS: "linux", ARCH: "armv7"}:
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
		dryRun    bool
		tags      string
		ldflags   string
		bingo     *flag.FlagSet = flag.NewFlagSet("bingo", flag.ContinueOnError)
	)

	bingo.StringVar(&Path, "d", Path, "output directory")
	bingo.StringVar(&tags, "tags", "", "build tags")
	bingo.StringVar(&ldflags, "ldflags", "", "ldflags")
	bingo.BoolVar(&stripFlag, "strip", false, "strip binary")
	bingo.BoolVar(&upxFlag, "upx", false, "compress binary with upx")
	bingo.BoolVar(&dryRun, "dry", false, "dry run")
	bingo.Parse(os.Args[1:])

	combos := []Combo{}
	freeArgs := []string{}

	log.Println("Parsing positional arguments:")
	for _, arg := range bingo.Args() {
		if !strings.Contains(arg, "/") {
			freeArgs = append(freeArgs, arg)
			log.Println("-", arg, "# passthrough")
			continue
		}
		log.Println("-", arg, "# combo")
		combos = append(combos, parseCombo(arg))
	}

	if len(combos) == 0 {
		combos = append(combos, DefaultCombo)
	}

	for _, c := range combos {
		fonts.Figlet(c.String())

		log.Println("Go Build Env:", c.Env())

		if dryRun {
			continue
		}

		buildArgs := []string{"build",
			"-o", filepath.Join(Path, c.ReleaseName()),
			"-mod=readonly",
			"-trimpath",
			"-ldflags", ldflags,
			"-tags", tags,
			"-v",
		}
		buildArgs = append(buildArgs, freeArgs...)
		buildArgs = append(buildArgs, "./cmd/k0s")

		for i, arg := range buildArgs {
			log.Println(fmt.Sprintf("[%02d] = %q", i, arg))
		}

		gocmd, ok := os.LookupEnv("GO")
		if !ok {
			gocmd = "go"
		}
		var (
			build = exec.Command(gocmd, buildArgs...)
			strip = exec.Command(
				"strip",
				filepath.Join(Path, c.ReleaseName()),
			)
			upx = exec.Command(
				"upx",
				filepath.Join(Path, c.ReleaseName()),
			)
		)

		log.Println("Start compiling...")
		build.Env = append(os.Environ(), c.Env()...)
		build.Stdout = os.Stdout
		build.Stderr = ts(os.Stderr)
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
			switch c.ARCH {
			case "arm64", "mips64", "mips64le", "s390x":
				break
			default:
				upx.Stdout = os.Stdout
				upx.Stderr = os.Stderr
				if err := upx.Run(); err != nil {
					log.Fatalln(err)
				}
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

func ts(next io.Writer) io.Writer {
	var (
		pr, pw  = io.Pipe()
		scanner = bufio.NewScanner(pr)
		logger  = log.New(next, "", log.Ldate|log.Ltime)
	)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			logger.Println(line)
		}
	}()

	return pw
}
