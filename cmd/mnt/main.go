package mnt

import (
	"log"
	"os"

	"k0s.io/pkg/cli/mnt"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	mnt.Run(os.Args[1:])
}
