module k0s.io/pkg/distro

go 1.19

require gitlab.com/mjwhitta/sysinfo v1.4.7

require (
	github.com/stretchr/testify v1.8.0 // indirect
	gitlab.com/mjwhitta/errors v1.0.0 // indirect
	gitlab.com/mjwhitta/hilighter v1.11.0 // indirect
	gitlab.com/mjwhitta/pathname v1.2.0 // indirect
	gitlab.com/mjwhitta/safety v1.11.0 // indirect
	gitlab.com/mjwhitta/where v1.2.4 // indirect
	golang.org/x/sys v0.0.0-20220712014510-0a85c31ab51e // indirect
)

replace k0s.io => ../../

replace k0s.io/cmd => ../../cmd/

replace k0s.io/pkg/agent => ../agent/

replace k0s.io/pkg/asciitransport => ../asciitransport/

replace k0s.io/pkg/cli => ../cli/

replace k0s.io/pkg/client => ../client/

replace k0s.io/pkg/console => ../console/

replace k0s.io/pkg/distro => ./

replace k0s.io/pkg/dohserver => ../dohserver/

replace k0s.io/pkg/exporter => ../exporter/

replace k0s.io/pkg/fonts => ../fonts/

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

replace k0s.io/pkg/ui => ../ui/
