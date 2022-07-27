module k0s.io/pkg/middleware

go 1.19

require (
	github.com/gorilla/handlers v1.5.1
	github.com/rs/cors v1.8.2
)

require github.com/felixge/httpsnoop v1.0.2 // indirect

replace k0s.io => ../../

replace k0s.io/cmd => ../../cmd/

replace k0s.io/pkg/agent => ../agent/

replace k0s.io/pkg/asciitransport => ../asciitransport/

replace k0s.io/pkg/cli => ../cli/

replace k0s.io/pkg/client => ../client/

replace k0s.io/pkg/console => ../console/

replace k0s.io/pkg/distro => ../distro/

replace k0s.io/pkg/dohserver => ../dohserver/

replace k0s.io/pkg/exporter => ../exporter/

replace k0s.io/pkg/fonts => ../fonts/

replace k0s.io/pkg/fzf => ../fzf/

replace k0s.io/pkg/gitd => ../gitd/

replace k0s.io/pkg/hub => ../hub/

replace k0s.io/pkg/jsondns => ../jsondns/

replace k0s.io/pkg/manager => ../manager/

replace k0s.io/pkg/middleware => ./

replace k0s.io/pkg/reverseproxy => ../reverseproxy/

replace k0s.io/pkg/rng => ../rng/

replace k0s.io/pkg/simple => ../simple/

replace k0s.io/pkg/tunnel => ../tunnel/

replace k0s.io/pkg/utils => ../utils/

replace k0s.io/pkg/uuid => ../uuid/

replace k0s.io/pkg/wrap => ../wrap/

replace k0s.io/pkg/ui => ../ui/

replace k0s.io/pkg/plugin => ../plugin/
