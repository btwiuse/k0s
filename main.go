package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
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

func usage(){
	fmt.Println(`please specify one of the subcommands: 
- agent
- hub
- client`)
}
