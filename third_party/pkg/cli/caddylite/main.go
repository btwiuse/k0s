package caddy

import (
	"os"

	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// plug in Caddy modules here
	_ "github.com/caddyserver/caddy/v2/modules/standard"
	_ "github.com/caddyserver/forwardproxy"
	_ "github.com/casbin/caddy-authz/v2"

	// _ "github.com/greenpau/caddy-auth-portal"
	_ "github.com/caddy-dns/alidns"
	_ "github.com/caddy-dns/azure"
	_ "github.com/caddy-dns/cloudflare"
	_ "github.com/caddy-dns/digitalocean"
	_ "github.com/caddy-dns/dnspod"
	_ "github.com/caddy-dns/duckdns"
	_ "github.com/caddy-dns/route53"
	_ "github.com/caddy-dns/vultr"
	_ "github.com/greenpau/caddy-auth-jwt"
	_ "github.com/greenpau/caddy-trace"
	_ "github.com/mholt/caddy-l4"
	_ "k0s.io/pkg/plugin/agent"
	_ "k0s.io/third_party/pkg/module/hub"
	_ "k0s.io/third_party/pkg/plugin/hello"
)

func Run(args []string) error {
	os.Args = append([]string{"caddy"}, args...)
	caddycmd.Main()
	return nil
}
