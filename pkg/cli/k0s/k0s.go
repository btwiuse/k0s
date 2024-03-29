package k0s

import (
	"fmt"
	"os"

	"k0s.io/pkg/cli/agent"
	"k0s.io/pkg/cli/chassis"
	"k0s.io/pkg/cli/client"
	"k0s.io/pkg/cli/hub"

	// "k0s.io/pkg/cli/kubectl"
	"k0s.io/pkg/cli/kuber"
	"k0s.io/pkg/cli/miniclient"
	"k0s.io/pkg/cli/mnt"
	"k0s.io/pkg/cli/upgrade"
	"k0s.io/pkg/cli/version"

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
	// "kubectl":    kubectl.Run,
	"kuber":   kuber.Run,
	"knot":    TODO,
	"version": version.Run,
}

func Run(args []string) error {
	return cmdRun.Run(os.Args[1:])
}
