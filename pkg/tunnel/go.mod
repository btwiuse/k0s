module k0s.io/pkg/tunnel

go 1.19

require (
	github.com/btwiuse/gost v0.0.1
	github.com/gorilla/handlers v1.5.1
	github.com/rs/cors v1.8.2
	k0s.io v0.0.0-00010101000000-000000000000
	k0s.io/pkg/wrap v0.1.4
	nhooyr.io/websocket v1.8.7
)

require (
	git.torproject.org/pluggable-transports/goptlib.git v1.2.0 // indirect
	github.com/LiamHaworth/go-tproxy v0.0.0-20190726054950-ef7efd7f24ed // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/coreos/go-iptables v0.6.0 // indirect
	github.com/dchest/siphash v1.2.2 // indirect
	github.com/docker/libcontainer v2.2.1+incompatible // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-gost/gosocks4 v0.0.1 // indirect
	github.com/go-gost/gosocks5 v0.3.0 // indirect
	github.com/go-gost/relay v0.1.1-0.20211123134818-8ef7fd81ffd7 // indirect
	github.com/go-gost/tls-dissector v0.0.2-0.20220408131628-aac992c27451 // indirect
	github.com/go-log/log v0.2.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.11 // indirect
	github.com/klauspost/reedsolomon v1.9.15 // indirect
	github.com/lucas-clemente/quic-go v0.28.0 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.5 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-18 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-19 v0.1.0-beta.1 // indirect
	github.com/miekg/dns v1.1.50 // indirect
	github.com/milosgajdos/tenus v0.0.3 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/riobard/go-bloom v0.0.0-20200614022211-cdc8013cb5b3 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/shadowsocks/go-shadowsocks2 v0.1.5 // indirect
	github.com/shadowsocks/shadowsocks-go v0.0.0-20200409064450-3e585ff90601 // indirect
	github.com/songgao/water v0.0.0-20200317203138-2b4b6d7c09d8 // indirect
	github.com/templexxx/cpu v0.0.7 // indirect
	github.com/templexxx/xorsimd v0.4.1 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/xtaci/kcp-go/v5 v5.6.1 // indirect
	github.com/xtaci/smux v1.5.16 // indirect
	github.com/xtaci/tcpraw v1.2.25 // indirect
	gitlab.com/yawning/obfs4.git v0.0.0-20210511220700-e330d1b7024b // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20220708220712-1185a9018129 // indirect
	golang.org/x/sys v0.0.0-20220712014510-0a85c31ab51e // indirect
	golang.org/x/text v0.3.8-0.20211004125949-5bd84dd9b33b // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
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

replace k0s.io/pkg/fzf => ../fzf/

replace k0s.io/pkg/gitd => ../gitd/

replace k0s.io/pkg/hub => ../hub/

replace k0s.io/pkg/jsondns => ../jsondns/

replace k0s.io/pkg/manager => ../manager/

replace k0s.io/pkg/middleware => ../middleware/

replace k0s.io/pkg/reverseproxy => ../reverseproxy/

replace k0s.io/pkg/rng => ../rng/

replace k0s.io/pkg/simple => ../simple/

replace k0s.io/pkg/tunnel => ./

replace k0s.io/pkg/utils => ../utils/

replace k0s.io/pkg/uuid => ../uuid/

replace k0s.io/pkg/wrap => ../wrap/

replace k0s.io/pkg/ui => ../ui/

replace k0s.io/pkg/plugin => ../plugin/
