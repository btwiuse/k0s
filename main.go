package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	match "github.com/alexpantyukhin/go-pattern-match"

	"k0s.io/pkg/cli/agent"
	"k0s.io/pkg/cli/bcrypt"
	"k0s.io/pkg/cli/buildkite"
	"k0s.io/pkg/cli/caddy"
	"k0s.io/pkg/cli/cadvisor"
	"k0s.io/pkg/cli/chassis"
	"k0s.io/pkg/cli/client"
	"k0s.io/pkg/cli/dohserver"
	"k0s.io/pkg/cli/filebrowser"
	"k0s.io/pkg/cli/gitd"
	"k0s.io/pkg/cli/goproxy"
	"k0s.io/pkg/cli/gos"
	"k0s.io/pkg/cli/gost"
	"k0s.io/pkg/cli/hub"
	"k0s.io/pkg/cli/k16s"
	"k0s.io/pkg/cli/kubectl"
	"k0s.io/pkg/cli/mnt"
	"k0s.io/pkg/cli/trojan"
	"k0s.io/pkg/cli/webproc"
)

var cmdRun = map[string]func([]string) error{
	"dohserver":       dohserver.Run,
	"bcrypt":          bcrypt.Run,
	"k16s":            k16s.Run,
	"kubectl":         kubectl.Run,
	"mnt":             mnt.Run,
	"webproc":         webproc.Run,
	"trojan":          trojan.Run,
	"goproxy":         goproxy.Run,
	"gos":             gos.Run,
	"buildkite-agent": buildkite.Run,
	"caddy":           caddy.Run,
	"cadvisor":        cadvisor.Run,
	"chassis":         chassis.Run,
	"client":          client.Run,
	"hub":             hub.Run,
	"hub2":            hub.Run2,
	"agent":           agent.Run,
	"gitd":            gitd.Run,
	"gost":            gost.Main,
	"filebrowser":     filebrowser.Run,
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	exe := strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")

	osargs := append([]string{exe}, os.Args[1:]...)

	// arg parse using rust-style match
	// https://github.com/ylxdzsw/v2socks/blob/master/src/main.rs
	// https://github.com/alexpantyukhin/go-pattern-match
	matcher := match.Match(osargs)

	for cmd := range cmdRun {
		subcmd := cmd
		runf, _ := cmdRun[cmd]
		// log.Println(subcmd)
		matcher = matcher.
			When(
				[]interface{}{
					subcmd,
					match.ANY,
				},
				func() error {
					// log.Println(subcmd)
					return runf(osargs[1:])
				},
			).
			When(
				[]interface{}{
					match.ANY,
					subcmd,
					match.ANY,
				},
				func() error {
					// log.Println(subcmd)
					return runf(osargs[2:])
				},
			)
	}

	matcher = matcher.When(
		[]interface{}{
			"k0s",
			match.ANY,
		},
		func() error {
			// log.Println("k0s")
			return client.Run(osargs[1:])
		},
	)

	ok, err := matcher.Result()
	if !ok {
		usage()
	}
	if err != nil {
		log.Fatalln(err)
	}
}

func usage() {
	fmt.Println("please specify one of the subcommands:")
	for c := range cmdRun {
		fmt.Println("-", c)
	}
}
