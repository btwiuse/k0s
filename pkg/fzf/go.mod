module k0s.io/pkg/fzf

go 1.21

require (
	github.com/gdamore/tcell v1.4.0
	github.com/mattn/go-isatty v0.0.14
	github.com/mattn/go-runewidth v0.0.13
	github.com/mattn/go-shellwords v1.0.12
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d
)

require (
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.0.3 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/term v0.0.0-20220411215600-e5f449aeb171 // indirect
	golang.org/x/text v0.3.8 // indirect
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

replace k0s.io/pkg/fonts => ../fonts/

replace k0s.io/pkg/fzf => ./

replace k0s.io/pkg/gitd => ../gitd/

replace k0s.io/pkg/hub => ../hub/

replace k0s.io/pkg/jsondns => ../jsondns/

replace k0s.io/pkg/manager => ../manager/

replace k0s.io/pkg/middleware => ../middleware/

replace k0s.io/pkg/plugin => ../plugin/

replace k0s.io/pkg/reverseproxy => ../reverseproxy/

replace k0s.io/pkg/rng => ../rng/

replace k0s.io/pkg/simple => ../simple/

replace k0s.io/pkg/tunnel => ../tunnel/

replace k0s.io/pkg/ui => ../ui/

replace k0s.io/pkg/utils => ../utils/

replace k0s.io/pkg/uuid => ../uuid/

replace k0s.io/pkg/wrap => ../wrap/

replace k0s.io/third_party => ../../third_party/

replace k0s.io/pkg/dial => ../dial/
