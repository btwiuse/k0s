package hub

import (
	"flag"
	"fmt"

	"github.com/caddyserver/caddy/v2"
	caddycmd "github.com/caddyserver/caddy/v2/cmd"
)

func init() {
	caddycmd.RegisterCommand(caddycmd.Command{
		Name:  "hub",
		Func:  cmdHub,
		Short: "hub",
		Long:  "hub",
		Flags: func() *flag.FlagSet {
			fset := flag.NewFlagSet("hub", flag.ExitOnError)
			fset.String("port", ":8000", "hub listening port")
			fset.String("cert", "", "path to tls cert file")
			fset.String("key", "", "path to tls key file")
			fset.Bool("version", false, "Show hub version info.")
			return fset
		}(),
	})
}

func cmdHub(fs caddycmd.Flags) (int, error) {
	fmt.Println("TODO")
	return caddy.ExitCodeSuccess, nil
}
