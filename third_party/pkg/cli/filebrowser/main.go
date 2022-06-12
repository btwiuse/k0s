package filebrowser

import (
	"os"
	"runtime"

	"github.com/filebrowser/filebrowser/v2/cmd"
)

func Run(args []string) error {
	os.Args = append([]string{"filebrowser"}, args...)
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
	return nil
}
