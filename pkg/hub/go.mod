module k0s.io/pkg/hub

go 1.19

replace (
	github.com/ginuerzh/gost => github.com/btwiuse/gost v0.0.0-20220706230444-f7c58a8aecbb
	k0s.io => ../../
	k0s.io/pkg/agent => ../agent
	k0s.io/pkg/api => ../api
	k0s.io/pkg/asciitransport => ../asciitransport
	k0s.io/pkg/distro => ../distro
	k0s.io/pkg/middleware => ../middleware
	k0s.io/pkg/rng => ../rng
	k0s.io/pkg/uuid => ../uuid
	k0s.io/pkg/version => ../version
	k0s.io/pkg/wrap => ../wrap
)

require (
	github.com/abbot/go-http-auth v0.4.0
	github.com/btwiuse/pretty v0.0.0-20220605073556-820155689895
	github.com/btwiuse/wetty v0.0.36
	github.com/gorilla/mux v1.8.0
	github.com/jpillora/go-echo-server v0.5.0
	k0s.io v0.0.0-00010101000000-000000000000
	k0s.io/pkg/agent v0.0.0-00010101000000-000000000000
	k0s.io/pkg/api v0.0.0-00010101000000-000000000000
	k0s.io/pkg/exporter v0.0.0-20220614032006-56adefb90b87
	k0s.io/pkg/middleware v0.0.0-00010101000000-000000000000
	k0s.io/pkg/rng v0.0.0-00010101000000-000000000000
	k0s.io/pkg/uuid v0.0.0-00010101000000-000000000000
	k0s.io/pkg/version v0.0.0-00010101000000-000000000000
	k0s.io/pkg/wrap v0.0.0-00010101000000-000000000000
	modernc.org/httpfs v1.0.6
	nhooyr.io/websocket v1.8.7
)

require (
	git.torproject.org/pluggable-transports/goptlib.git v1.2.0 // indirect
	github.com/ActiveState/termtest/conpty v0.5.0 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/LiamHaworth/go-tproxy v0.0.0-20190726054950-ef7efd7f24ed // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15 // indirect
	github.com/andrew-d/go-termutil v0.0.0-20150726205930-009166a695a2 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/beevik/ntp v0.3.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/btwiuse/gods v0.0.0-20190414062120-7e7cf0aebbb0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/coreos/go-iptables v0.6.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/creack/pty v1.1.18 // indirect
	github.com/dchest/siphash v1.2.2 // indirect
	github.com/denisbrodbeck/machineid v1.0.1 // indirect
	github.com/docker/docker v20.10.17+incompatible // indirect
	github.com/docker/libcontainer v2.2.1+incompatible // indirect
	github.com/ema/qdisc v0.0.0-20200603082823-62d0308e3e00 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/ginuerzh/gost v0.0.0-20210206051340-8dd4d8d9a123 // indirect
	github.com/go-gost/gosocks4 v0.0.1 // indirect
	github.com/go-gost/gosocks5 v0.3.0 // indirect
	github.com/go-gost/relay v0.1.1-0.20211123134818-8ef7fd81ffd7 // indirect
	github.com/go-gost/tls-dissector v0.0.2-0.20220408131628-aac992c27451 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-log/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/godbus/dbus v0.0.0-20190422162347-ade71ed3457e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/go-envparse v0.0.0-20200406174449-d9cfd743a15e // indirect
	github.com/hodgesds/perf-utils v0.4.0 // indirect
	github.com/illumos/go-kstat v0.0.0-20210513183136-173c9b0a9973 // indirect
	github.com/josharian/native v0.0.0-20200817173448-b6b71def0850 // indirect
	github.com/jpillora/ansi v1.0.2 // indirect
	github.com/jpillora/requestlog v1.0.0 // indirect
	github.com/jpillora/sizestr v1.0.0 // indirect
	github.com/jsimonetti/rtnetlink v0.0.0-20211022192332-93da33804786 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.11 // indirect
	github.com/klauspost/reedsolomon v1.9.15 // indirect
	github.com/lucas-clemente/quic-go v0.28.0 // indirect
	github.com/lufia/iostat v1.2.0 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.5 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-18 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-19 v0.1.0-beta.1 // indirect
	github.com/mattn/go-xmlrpc v0.0.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mdlayher/genetlink v1.0.0 // indirect
	github.com/mdlayher/netlink v1.4.1 // indirect
	github.com/mdlayher/socket v0.0.0-20210307095302-262dc9984e00 // indirect
	github.com/mdlayher/wifi v0.0.0-20200527114002-84f0b9457fdd // indirect
	github.com/miekg/dns v1.1.48 // indirect
	github.com/milosgajdos/tenus v0.0.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.34.0 // indirect
	github.com/prometheus/node_exporter v1.3.1 // indirect
	github.com/prometheus/procfs v0.7.4-0.20211011103944-1a7a2bd3279f // indirect
	github.com/riobard/go-bloom v0.0.0-20200614022211-cdc8013cb5b3 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/safchain/ethtool v0.1.0 // indirect
	github.com/shadowsocks/go-shadowsocks2 v0.1.5 // indirect
	github.com/shadowsocks/shadowsocks-go v0.0.0-20200409064450-3e585ff90601 // indirect
	github.com/songgao/water v0.0.0-20200317203138-2b4b6d7c09d8 // indirect
	github.com/soundcloud/go-runit v0.0.0-20150630195641-06ad41a06c4a // indirect
	github.com/templexxx/cpu v0.0.7 // indirect
	github.com/templexxx/xorsimd v0.4.1 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce // indirect
	github.com/xtaci/kcp-go/v5 v5.6.1 // indirect
	github.com/xtaci/smux v1.5.16 // indirect
	github.com/xtaci/tcpraw v1.2.25 // indirect
	gitlab.com/mjwhitta/errors v1.0.0 // indirect
	gitlab.com/mjwhitta/hilighter v1.11.0 // indirect
	gitlab.com/mjwhitta/pathname v1.2.0 // indirect
	gitlab.com/mjwhitta/safety v1.11.0 // indirect
	gitlab.com/mjwhitta/sysinfo v1.4.7 // indirect
	gitlab.com/mjwhitta/where v1.2.4 // indirect
	gitlab.com/yawning/obfs4.git v0.0.0-20210511220700-e330d1b7024b // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	golang.org/x/sys v0.0.0-20220708085239-5a0f0661e09d // indirect
	golang.org/x/text v0.3.8-0.20211004125949-5bd84dd9b33b // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	google.golang.org/genproto v0.0.0-20220706185917-7780775163c4 // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k0s.io/pkg/asciitransport v0.0.0-00010101000000-000000000000 // indirect
	k0s.io/pkg/distro v0.0.0-00010101000000-000000000000 // indirect
)
