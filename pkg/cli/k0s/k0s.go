package k0s

import (
	"fmt"
	"os"

	"k0s.io/pkg/cli/agent"
	"k0s.io/pkg/cli/chassis"
	"k0s.io/pkg/cli/client"
	"k0s.io/pkg/cli/hub"
	"k0s.io/pkg/cli/miniclient"
	"k0s.io/pkg/cli/mnt"
	"k0s.io/pkg/cli/upgrade"

	"github.com/btwiuse/multicall"
)

func TODO([]string) error {
	fmt.Println("TODO: not implemented yet")
	return nil
}

var cmdRun multicall.RunnerFuncMap = map[string]multicall.RunnerFunc{
	"mnt":        mnt.Run,
	"chassis":    chassis.Run,
	"client":     client.Run,
	"miniclient": miniclient.Run,
	"hub":        hub.Run,
	"hub2":       hub.Run2,
	"agent":      agent.Run,
	"upgrade":    upgrade.Run,
	"kuber":      TODO,
	"knot":       TODO,
	"version":    TODO,
}

func Run(args []string) error {
	return cmdRun.Run(os.Args[1:])
}
