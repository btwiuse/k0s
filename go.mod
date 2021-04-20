module k0s.io

go 1.16

replace (
	github.com/buildkite/agent/v3 => github.com/btwiuse/agent/v3 v3.27.1-0.20210217080418-ae42a28eefa7
	github.com/caddyserver/forwardproxy => github.com/klzgrad/forwardproxy v0.0.0-20210120121422-9b4a5a242bd6
	github.com/coredns/coredns => github.com/btwiuse/coredns v1.8.4
	github.com/ginuerzh/gost => github.com/btwiuse/gost v0.0.0-20210218034515-4e5ef1691e9f
	github.com/greenpau/caddy-auth-portal => github.com/btwiuse/caddy-auth-portal v1.3.12-0.20210204101408-068c2618b417
	k8s.io/api => k8s.io/api v0.20.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.4
	k8s.io/apiserver => github.com/btwiuse/apiserver v0.20.4-btwiuse
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.4
	k8s.io/client-go => k8s.io/client-go v0.20.4
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.4
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.4
	k8s.io/code-generator => k8s.io/code-generator v0.20.4
	k8s.io/component-base => k8s.io/component-base v0.20.4
	k8s.io/component-helpers => k8s.io/component-helpers v0.20.4
	k8s.io/controller-manager => k8s.io/controller-manager v0.20.4
	k8s.io/cri-api => k8s.io/cri-api v0.20.4
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.4
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.4
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.4
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.4
	k8s.io/kube-state-metrics/v2 => github.com/btwiuse/k16s/v2 v2.0.0-beta.0.20201224174453-2114e07844a9
	k8s.io/kubectl => k8s.io/kubectl v0.20.4
	k8s.io/kubelet => k8s.io/kubelet v0.20.4
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.4
	k8s.io/metrics => k8s.io/metrics v0.20.4
	k8s.io/mount-utils => k8s.io/mount-utils v0.20.4
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.20.4
	nhooyr.io/websocket => github.com/btwiuse/websocket v1.8.6
)

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/MaxRis/w32 v0.0.0-20180517000239-4f5cfb03fabf // indirect
	github.com/VojtechVitek/yaml-cli v0.0.5
	github.com/abbot/go-http-auth v0.4.0
	github.com/abiosoft/caddy-json-schema v0.0.0-20200527180432-2d0cb96ed8ea
	github.com/abiosoft/ishell v2.0.0+incompatible
	github.com/abiosoft/readline v0.0.0-20180607040430-155bce2042db // indirect
	github.com/adrg/xdg v0.3.3
	github.com/alexpantyukhin/go-pattern-match v0.0.0-20200628201436-c57d5ad3f2c5
	github.com/btwiuse/bcrypt v1.0.2
	github.com/btwiuse/cadvisor v0.0.0-20210312172035-34fddda41018
	github.com/btwiuse/dkg v0.2.0
	github.com/btwiuse/etcd/v3 v3.4.15
	github.com/btwiuse/gods v0.0.0-20190414062120-7e7cf0aebbb0
	github.com/btwiuse/pretty v0.0.0-20190401073227-519ff4ea1882
	github.com/btwiuse/wetty v0.0.36
	github.com/buildkite/agent/v3 v3.27.0
	github.com/caddy-dns/alidns v1.0.21
	github.com/caddy-dns/azure v0.2.0
	github.com/caddy-dns/cloudflare v0.0.0-20210401224357-964e47d3890e
	github.com/caddy-dns/digitalocean v0.0.0-20210408173619-385f9346b5ac
	github.com/caddy-dns/dnspod v0.0.1
	github.com/caddy-dns/duckdns v0.3.1
	github.com/caddy-dns/route53 v1.1.1
	github.com/caddy-dns/vultr v0.0.0-20210105121117-3162aa6b9c27
	github.com/caddyserver/caddy/v2 v2.4.0-beta.2
	github.com/caddyserver/forwardproxy v0.0.0-00010101000000-000000000000
	github.com/casbin/caddy-authz/v2 v2.0.0
	github.com/coredns/coredns v1.8.4-0.20210224180316-9d3a84377cae
	github.com/cppforlife/go-cli-ui v0.0.0-20200716203538-1e47f820817f
	github.com/creack/pty v1.1.11
	github.com/dchest/siphash v1.2.1 // indirect
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/docker/docker v20.10.2+incompatible
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/filebrowser/filebrowser/v2 v2.11.0
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/freman/caddy2-reauth v0.0.0-20200518130136-6064aa96b1a8
	github.com/gdamore/tcell v1.3.0
	github.com/ginuerzh/gost v0.0.0-20210206051340-8dd4d8d9a123
	github.com/go-log/log v0.2.0
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/google/cadvisor v0.39.0
	github.com/google/uuid v1.2.0
	github.com/goproxyio/goproxy/v2 v2.0.5
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/greenpau/caddy-auth-jwt v1.2.6
	github.com/greenpau/caddy-auth-portal v0.0.0-00010101000000-000000000000
	github.com/greenpau/caddy-trace v1.1.5
	github.com/hairyhenderson/caddy-teapot-module v0.0.2
	github.com/iamd3vil/caddy_yaml_adapter v0.0.0-20200503183711-d479c29b475a
	github.com/infobloxopen/go-trees v0.0.0-20200715205103-96a057b8dfb9
	github.com/ipfs/go-cid v0.0.7
	github.com/jaypipes/ghw v0.7.0 // indirect
	github.com/jpillora/chisel v1.7.6 // indirect
	github.com/jpillora/opts v1.2.0
	github.com/jpillora/webproc v0.4.0
	github.com/json-iterator/go v1.1.10
	github.com/klauspost/cpuid v1.3.1 // indirect
	github.com/koding/websocketproxy v0.0.0-20181220232114-7ed82d81a28c // indirect
	github.com/liamg/aminal v0.9.0
	github.com/libp2p/go-libp2p v0.13.0
	github.com/libp2p/go-libp2p-core v0.8.5
	github.com/libp2p/go-libp2p-kad-dht v0.11.1
	github.com/lucasb-eyer/go-colorful v1.0.3 // indirect
	github.com/lukesampson/figlet v0.0.0-20190211215653-8a3ef4a6ac42
	github.com/lunny/tango v0.5.6
	github.com/mathetake/gasm v0.0.0-20200928142744-80e74517647c
	github.com/mattn/go-isatty v0.0.12
	github.com/mattn/go-runewidth v0.0.10
	github.com/mattn/go-shellwords v1.0.10
	github.com/mholt/caddy-l4 v0.0.0-20210209073014-d1d54b015e34
	github.com/mholt/caddy-webdav v0.0.0-20200916200058-c949b3226234
	github.com/miekg/dns v1.1.40
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/multiformats/go-multihash v0.0.15
	github.com/multiformats/go-varint v0.0.6
	github.com/p4gefau1t/trojan-go v0.8.2
	github.com/pkg/errors v0.9.1
	github.com/portainer/agent v0.0.0-20210129020346-65fa249df79a
	github.com/portainer/libcrypto v0.0.0-20201216185936-7a703a1ea452 // indirect
	github.com/portainer/libhttp v0.0.0-20201216185909-d20481a3da82 // indirect
	github.com/prometheus/client_golang v1.9.0
	github.com/prometheus/common v0.15.0
	github.com/prometheus/node_exporter v1.0.1
	github.com/pupapaik/sysinfo v0.0.0-20200106202926-c17dea004cd4
	github.com/rancher/dapper v0.5.5
	github.com/rs/cors v1.7.0
	github.com/schollz/pake/v2 v2.0.7
	github.com/schollz/progressbar/v3 v3.8.0
	github.com/sirupsen/logrus v1.8.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tango-contrib/basicauth v0.0.0-20170526072948-7fbc19aece86
	github.com/traefik/yaegi v0.9.14
	github.com/tyler-smith/go-bip39 v1.1.0
	github.com/urfave/cli v1.22.5
	github.com/urfave/cli/v2 v2.3.0
	github.com/v2fly/v2ray-core/v4 v4.36.2
	github.com/whyrusleeping/mdns v0.0.0-20190826153040-b9b60ed33aa9
	go.starlark.net v0.0.0-20210305151048-6a590ae7f4eb
	golang.org/x/crypto v0.0.0-20210314154223-e6e6c4f2bb5b
	golang.org/x/mod v0.4.1
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/apiserver v0.20.4
	k8s.io/autoscaler/vertical-pod-autoscaler v0.9.2
	k8s.io/client-go v0.20.4
	k8s.io/component-base v0.20.4
	k8s.io/klog/v2 v2.7.0
	k8s.io/kube-state-metrics/v2 v2.0.0-00010101000000-000000000000
	k8s.io/kubectl v0.20.4
	k8s.io/kubelet v0.0.0 // indirect
	modernc.org/httpfs v1.0.0
	nhooyr.io/websocket v1.8.6
	robpike.io/ivy v0.1.0
	sigs.k8s.io/metrics-server v0.4.2
	src.elv.sh v0.14.1-0.20210218105754-53593c3ab79f
)
