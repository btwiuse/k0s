package mnt

import (
	"os"

	"k0s.io/pkg/cli/mnt"
)

func main() {
	mnt.Run(os.Args[1:])
}
