package jsonschema

import (
	"flag"
	"fmt"
	stdlog "log"
	"os"

	caddycmd "github.com/caddyserver/caddy/v2/cmd"
)

const commandName = "hello"

var (
	config = struct {
		VsCode bool
	}{}

	log = stdlog.New(os.Stderr, commandName+" ", 0)
)

func init() {
	caddycmd.RegisterCommand(caddycmd.Command{
		Name:  commandName,
		Func:  run,
		Usage: "[--vscode]",
		Short: "Say hello",
		Long: `
Say hello

If --vscode is set, schema and vscode config is generated into a '.vscode' directory
in the current working directory. This disregards '--output'.
Other ways of integrating JSON schema in VSCode can be found at
https://code.visualstudio.com/docs/languages/json#_mapping-in-the-user-settings
`,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet(commandName, flag.ExitOnError)
			fs.BoolVar(&config.VsCode, "vscode", config.VsCode, "Generate VSCode configuration")
			return fs
		}(),
	})
}

func run(fs caddycmd.Flags) (int, error) {
	w := "vscode"
	if config.VsCode {
		w = "true"
	} else {
		w = "false"
	}

	fmt.Println("hello", w)

	return 0, nil
}
