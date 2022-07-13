module k0s.io/pkg/fonts

go 1.19

require (
	github.com/lukesampson/figlet v0.0.0-20190211215653-8a3ef4a6ac42
	github.com/mattn/go-isatty v0.0.14
	k0s.io/pkg/console v0.0.0-00010101000000-000000000000
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.0.0-20220712014510-0a85c31ab51e // indirect
)

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

replace k0s.io/pkg/fonts => ./

replace k0s.io/pkg/fzf => ../fzf/

replace k0s.io/pkg/gitd => ../gitd/

replace k0s.io/pkg/hub => ../hub/

replace k0s.io/pkg/jsondns => ../jsondns/

replace k0s.io/pkg/manager => ../manager/

replace k0s.io/pkg/middleware => ../middleware/

replace k0s.io/pkg/reverseproxy => ../reverseproxy/

replace k0s.io/pkg/rng => ../rng/

replace k0s.io/pkg/simple => ../simple/

replace k0s.io/pkg/tunnel => ../tunnel/

replace k0s.io/pkg/utils => ../utils/

replace k0s.io/pkg/uuid => ../uuid/

replace k0s.io/pkg/wrap => ../wrap/
