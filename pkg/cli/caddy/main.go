package caddy

import (
        "os"
        caddycmd "github.com/caddyserver/caddy/v2/cmd"

        // plug in Caddy modules here
        _ "github.com/caddyserver/caddy/v2/modules/standard"
        _ "github.com/caddyserver/forwardproxy"
        _ "github.com/caddyserver/nginx-adapter"
        _ "github.com/casbin/caddy-authz/v2"
        _ "github.com/freman/caddy2-reauth"
        _ "github.com/greenpau/caddy-auth-portal"
        _ "github.com/greenpau/caddy-auth-jwt"
        _ "github.com/greenpau/caddy-trace"
        _ "github.com/hairyhenderson/caddy-teapot-module"
        _ "github.com/iamd3vil/caddy_yaml_adapter"
        _ "github.com/mholt/caddy-webdav"
        _ "github.com/mholt/caddy-l4"
	_ "github.com/abiosoft/caddy-json-schema"
	_ "k0s.io/pkg/plugin/agent"
	_ "k0s.io/pkg/plugin/hello"
)

func Run(args []string) error {
	os.Args = append([]string{"caddy"}, args...)
        caddycmd.Main()
	return nil
}
