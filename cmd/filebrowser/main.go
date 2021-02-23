package main

import (
	"os"
	"runtime"

	"k0s.io/pkg/cli/filebrowser"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	filebrowser.Run(os.Args[1:])
}
