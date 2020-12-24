package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexpantyukhin/go-pattern-match"

	"k0s.io/k0s/pkg/cli/agent"
	"k0s.io/k0s/pkg/cli/chassis"
	"k0s.io/k0s/pkg/cli/client"
	"k0s.io/k0s/pkg/cli/hub"
	"k0s.io/k0s/pkg/cli/gitd"
	"k0s.io/k0s/pkg/cli/gost"
	"k0s.io/k0s/pkg/cli/mnt"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	exe := strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")

	osargs := append([]string{exe}, os.Args[1:]...)

	// arg parse using rust-style match
	// https://github.com/ylxdzsw/v2socks/blob/master/src/main.rs
	// https://github.com/alexpantyukhin/go-pattern-match
	match.Match(osargs).

		// hub -> hub
		// agent -> agent
		// client -> client
		When([]interface{}{"mnt", match.ANY}, func() {
			log.Fatalln(mnt.Run(osargs[1:]))
		}).
		When([]interface{}{"chassis", match.ANY}, func() {
			log.Fatalln(chassis.Run(osargs[1:]))
		}).
		When([]interface{}{"client", match.ANY}, func() {
			log.Fatalln(client.Run(osargs[1:]))
		}).
		When([]interface{}{"hub", match.ANY}, func() {
			log.Fatalln(hub.Run(osargs[1:]))
		}).
		When([]interface{}{"hub2", match.ANY}, func() {
			log.Fatalln(hub.Run2(osargs[1:]))
		}).
		When([]interface{}{"agent", match.ANY}, func() {
			log.Fatalln(agent.Run(osargs[1:]))
		}).
		When([]interface{}{"gitd", match.ANY}, func() {
			log.Fatalln(gitd.Run(osargs[1:]))
		}).
		When([]interface{}{"gost", match.ANY}, func() {
			gost.Main(osargs[1:])
		}).

		// conntroll hub -> hub
		// conntroll agent -> agent
		// conntroll client -> client
		// k0s hub -> hub
		// k0s agent -> agent
		// k0s client -> client
		// * hub -> hub
		// * agent -> agent
		// * client -> client
		When([]interface{}{match.ANY, "mnt", match.ANY}, func() {
			log.Fatalln(mnt.Run(osargs[2:]))
		}).
		When([]interface{}{match.ANY, "chassis", match.ANY}, func() {
			log.Fatalln(chassis.Run(osargs[2:]))
		}).
		When([]interface{}{match.ANY, "client", match.ANY}, func() {
			log.Fatalln(client.Run(osargs[2:]))
		}).
		When([]interface{}{match.ANY, "hub", match.ANY}, func() {
			log.Fatalln(hub.Run(osargs[2:]))
		}).
		When([]interface{}{match.ANY, "hub2", match.ANY}, func() {
			log.Fatalln(hub.Run2(osargs[2:]))
		}).
		When([]interface{}{match.ANY, "agent", match.ANY}, func() {
			log.Fatalln(agent.Run(osargs[2:]))
		}).
		When([]interface{}{match.ANY, "gitd", match.ANY}, func() {
			log.Fatalln(gitd.Run(osargs[2:]))
		}).
		When([]interface{}{match.ANY, "gost", match.ANY}, func() {
			gost.Main(osargs[2:])
		}).

		// k0s -> client
		// k0s hub -> hub
		// k0s agent -> agent
		When([]interface{}{"k0s", match.ANY}, func() {
			log.Fatalln(client.Run(osargs[1:]))
		}).

		// conntroll -> usage
		When(match.ANY, usage).
		Result()
}

func usage() {
	fmt.Println(`please specify one of the subcommands: 
- agent
- hub
- client`)
}
