package main

import (
	"fmt"
	"log"
	"os"

	"github.com/btwiuse/multicall"
	"k0s.io/pkg/cli/agent"
	"k0s.io/pkg/cli/chassis"
	"k0s.io/pkg/cli/client"
	"k0s.io/pkg/cli/hub"
	"k0s.io/pkg/cli/miniclient"
	"k0s.io/pkg/cli/mnt"
	"k0s.io/third_party/pkg/cli/bcrypt"
	"k0s.io/third_party/pkg/cli/buildkite"

	// "k0s.io/third_party/pkg/cli/caddy"
	"k0s.io/pkg/cli/kubectl"
	"k0s.io/third_party/pkg/cli/dohserver"
	"k0s.io/third_party/pkg/cli/filebrowser"
	"k0s.io/third_party/pkg/cli/gitd"
	"k0s.io/third_party/pkg/cli/goproxy"
	"k0s.io/third_party/pkg/cli/gos"

	// "k0s.io/third_party/pkg/cli/gost"
	"k0s.io/third_party/pkg/cli/k16s"
	"k0s.io/third_party/pkg/cli/trojan"
	"k0s.io/third_party/pkg/cli/webproc"
)

func TODO([]string) error {
	fmt.Println("TODO: not implemented yet")
	return nil
}

var cmdRun multicall.RunnerFuncMap = map[string]multicall.RunnerFunc{
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
	"caddy":           TODO, // caddy.Run,
	"chassis":         chassis.Run,
	"client":          client.Run,
	"miniclient":      miniclient.Run,
	"hub":             hub.Run,
	"hub2":            hub.Run2,
	"agent":           agent.Run,
	"gitd":            gitd.Run,
	// "gost":            gost.Main,
	"filebrowser": filebrowser.Run,
}

func Run(args []string) error {
	return cmdRun.Run(os.Args[1:])
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err := Run(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}
