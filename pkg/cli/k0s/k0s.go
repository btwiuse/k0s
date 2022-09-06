package k0s

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	match "github.com/alexpantyukhin/go-pattern-match"

	"k0s.io/pkg/cli/agent"
	"k0s.io/pkg/cli/chassis"
	"k0s.io/pkg/cli/client"
	"k0s.io/pkg/cli/hub"
	"k0s.io/pkg/cli/miniclient"
	"k0s.io/pkg/cli/mnt"
	"k0s.io/pkg/cli/upgrade"
)

var cmdRun = map[string]func([]string) error{
	"mnt":        mnt.Run,
	"chassis":    chassis.Run,
	"client":     client.Run,
	"miniclient": miniclient.Run,
	"hub":        hub.Run,
	"hub2":       hub.Run2,
	"agent":      agent.Run,
	"upgrade":    upgrade.Run,
}

func Run(args []string) error {
	exe := strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")

	osargs := append([]string{exe}, args...)

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

	ok, _ := matcher.Result()
	if !ok {
		usage()
	}

	return nil
}

func usage() {
	fmt.Println("please specify one of the subcommands:")
	for c := range cmdRun {
		fmt.Println("-", c)
	}
}
