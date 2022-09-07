//go:build ignore
// +build ignore

package main

import (
	"os"
	"runtime"

	"k0s.io/third_party/pkg/cli/filebrowser"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	filebrowser.Run(os.Args[1:])
}
