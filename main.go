package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	exe, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}

	exe = filepath.Base(exe)

	switch {
	case strings.HasSuffix(exe, "agent"):
		agentCmd(os.Args[1:])
	case strings.HasSuffix(exe, "hub"):
		hubCmd(os.Args[1:])
	case strings.HasSuffix(exe, "client"):
		clientCmd(os.Args[1:])
	default:
		if len(os.Args) < 2 {
			usage()
			os.Exit(1)
		}
		switch subcmd := os.Args[1]; subcmd {
		case "agent":
			agentCmd(os.Args[2:])
		case "hub":
			hubCmd(os.Args[2:])
		case "client":
			clientCmd(os.Args[2:])
		default:
			log.Fatalln("unknown subcommand:", subcmd)
		}
	}
}

func usage() {
	fmt.Println(`please specify one of the subcommands: 
- agent
- hub
- client`)
}
