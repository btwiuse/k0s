package mnt

import (
	"os"

	"k0s.io/pkg/cli/mnt"
	"k0s.io/pkg/log"
)

func main() {
	mnt.Run(os.Args[1:])
}
