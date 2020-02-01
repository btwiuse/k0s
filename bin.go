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
)

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
	return strings.Join([]string{c.OS, c.ARCH, Basename}, Delimeter)
}

func (c Combo) Env() []string {
	envs := append(
		GlobalEnv,
		fmt.Sprintf("GOOS=%s", c.OS),
		fmt.Sprintf("GOARCH=%s", c.ARCH),
	)
	switch c.OS {
	case "android", "windows":
		envs = append(envs, "CGO_ENABLED=1")
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
			strip.Run()
		}

		if upxFlag {
			upx.Stdout = os.Stdout
			upx.Stderr = os.Stderr
			if err := upx.Run(); err != nil {
				log.Fatalln(err)
			}
		}

		if c == DefaultCombo {
			for _, bin := range []string{"k0s"} {
				src := filepath.Join(Path, c.ReleaseName())
				dst := filepath.Join(Path, bin)
				ln(src, dst)
			}
		}
	}
}

func ln(from, to string) {
	lnf := exec.Command("ln", "-f", "-v", from, to)
	lnf.Stdout = os.Stdout
	lnf.Stderr = os.Stderr
	if err := lnf.Run(); err != nil {
		log.Fatalln(err)
	}
}
