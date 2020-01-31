package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/alexpantyukhin/go-pattern-match"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	exe := filepath.Base(os.Args[0])

	osargs := append([]string{exe}, os.Args[1:]...)

	// arg parse using rust-style match
	// https://github.com/ylxdzsw/v2socks/blob/master/src/main.rs
	// https://github.com/alexpantyukhin/go-pattern-match
	match.Match(osargs).

		// hub -> hub
		// agent -> agent
		// client -> client
		When([]interface{}{"client", match.ANY}, func() {
			clientCmd(osargs[1:])
		}).
		When([]interface{}{"hub", match.ANY}, func() {
			hubCmd(osargs[1:])
		}).
		When([]interface{}{"agent", match.ANY}, func() {
			agentCmd(osargs[1:])
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
		When([]interface{}{match.ANY, "client", match.ANY}, func() {
			clientCmd(osargs[2:])
		}).
		When([]interface{}{match.ANY, "hub", match.ANY}, func() {
			hubCmd(osargs[2:])
		}).
		When([]interface{}{match.ANY, "agent", match.ANY}, func() {
			agentCmd(osargs[2:])
		}).

		// k0s -> client
		// k0s hub -> hub
		// k0s agent -> agent
		When([]interface{}{"k0s", match.ANY}, func() {
			clientCmd(osargs[1:])
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
