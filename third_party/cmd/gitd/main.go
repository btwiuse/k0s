package main

import (
	"log"
	"os"

	"k0s.io/third_party/pkg/cli/gitd"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Fatalln(gitd.Run(os.Args[1:]))
}
