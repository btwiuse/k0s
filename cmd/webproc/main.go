package main

import (
	"log"
	"os"

	"k0s.io/pkg/cli/webproc"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Fatalln(webproc.Run(os.Args[1:]))
}
