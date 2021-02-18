package caddy

import (
        "os"
        caddycmd "github.com/caddyserver/caddy/v2/cmd"

        // plug in Caddy modules here
        _ "github.com/caddyserver/caddy/v2/modules/standard"
        _ "github.com/greenpau/caddy-auth-portal"
        _ "github.com/greenpau/caddy-auth-jwt"
        _ "github.com/greenpau/caddy-trace"
        _ "github.com/caddyserver/forwardproxy"
        _ "github.com/caddyserver/nginx-adapter"
        _ "github.com/iamd3vil/caddy_yaml_adapter"
        _ "github.com/freman/caddy2-reauth"
        _ "github.com/mholt/caddy-webdav"
)

func Run(args []string) error {
	os.Args = append([]string{"caddy"}, args...)
        caddycmd.Main()
	return nil
}
