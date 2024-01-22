module k0s.io/third_party

go 1.21.6

replace github.com/buildkite/agent/v3 => github.com/btwiuse/agent/v3 v3.27.1-0.20210217080418-ae42a28eefa7

require (
	github.com/Dreamacro/clash v1.11.4
	github.com/alexpantyukhin/go-pattern-match v0.0.0-20200628201436-c57d5ad3f2c5 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr v1.4.10 // indirect
	github.com/btwiuse/bcrypt v1.0.2
	github.com/btwiuse/cadvisor v0.0.0-20210312172035-34fddda41018
	github.com/btwiuse/dkg v0.2.0
	github.com/btwiuse/gost v0.0.4
	github.com/btwiuse/wetty v0.0.36
	github.com/buildkite/agent/v3 v3.27.0
	github.com/caddy-dns/alidns v1.0.23
	github.com/caddy-dns/azure v0.2.0
	github.com/caddy-dns/cloudflare v0.0.0-20210607183747-91cf700356a1
	github.com/caddy-dns/digitalocean v0.0.0-20210809220558-ac6e4fd9e135
	github.com/caddy-dns/dnspod v0.0.4
	github.com/caddy-dns/duckdns v0.3.1
	github.com/caddy-dns/route53 v1.1.3
	github.com/caddy-dns/vultr v0.0.0-20211122185502-733392841379
	github.com/caddyserver/caddy/v2 v2.5.2
	github.com/caddyserver/forwardproxy v0.0.0-20211013034647-8c6ef2bd4a8f
	github.com/casbin/caddy-authz/v2 v2.0.0
	github.com/cppforlife/go-cli-ui v0.0.0-20200716203538-1e47f820817f
	github.com/creack/pty v1.1.18
	github.com/ethereum/go-ethereum v1.10.20
	github.com/filebrowser/filebrowser/v2 v2.11.0
	github.com/freman/caddy2-reauth v0.0.0-20200518130136-6064aa96b1a8
	github.com/go-log/log v0.2.0
	github.com/golang/protobuf v1.5.3
	github.com/google/cadvisor v0.39.0
	github.com/goproxyio/goproxy/v2 v2.0.5
	github.com/gorilla/handlers v1.5.2
	github.com/gorilla/rpc v1.2.0
	github.com/gorilla/websocket v1.5.0
	github.com/greenpau/caddy-auth-jwt v1.2.6
	github.com/greenpau/caddy-trace v1.1.8
	github.com/jpillora/opts v1.2.0
	github.com/jpillora/webproc v0.4.0
	github.com/lunny/tango v0.5.6
	github.com/mholt/caddy-l4 v0.0.0-20220420174601-aec6535658b1
	github.com/p4gefau1t/trojan-go v0.8.2
	github.com/portainer/agent v0.0.0-20220708094808-9c896747b4ff
	github.com/rancher/dapper v0.5.5
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.1
	github.com/tango-contrib/basicauth v0.0.0-20170526072948-7fbc19aece86
	github.com/tetratelabs/wazero v1.0.0-beta1
	github.com/traefik/yaegi v0.9.14
	github.com/urfave/cli v1.22.5
	github.com/v2fly/v2ray-core/v4 v4.36.2
	go.etcd.io/etcd/client/v3 v3.5.7
	go.starlark.net v0.0.0-20210305151048-6a590ae7f4eb
	go.uber.org/automaxprocs v1.5.1
	golang.org/x/crypto v0.18.0
	golang.org/x/mod v0.14.0
	golang.org/x/text v0.14.0
	google.golang.org/protobuf v1.28.1
	k0s.io v0.1.6
	k0s.io/pkg/agent v0.1.6
	k0s.io/pkg/asciitransport v0.1.6
	k0s.io/pkg/cli v0.1.5
	k0s.io/pkg/console v0.1.5
	k0s.io/pkg/dohserver v0.1.5
	k0s.io/pkg/exporter v0.1.6
	k0s.io/pkg/gitd v0.1.5
	k0s.io/pkg/plugin v0.1.5
	k0s.io/pkg/rng v0.1.5
	k0s.io/pkg/tunnel v0.1.5
	k0s.io/pkg/utils v0.1.5
	k0s.io/pkg/uuid v0.1.6
	// k8s.io/apimachinery v0.22.5
	// k8s.io/apiserver v0.22.5
	// k8s.io/client-go v0.22.5
	// k8s.io/component-base v0.22.5
	// k8s.io/klog/v2 v2.60.1
	// k8s.io/kubectl v0.20.4
	modernc.org/httpfs v1.0.6
	robpike.io/ivy v0.1.0
	// sigs.k8s.io/metrics-server v0.4.2
	src.elv.sh v0.14.1-0.20210218105754-53593c3ab79f
)

require (
	github.com/btwiuse/multicall v0.0.2
	k8s.io/api v0.27.3
	k8s.io/apimachinery v0.27.3
	k8s.io/client-go v0.27.3
	k8s.io/klog/v2 v2.90.1
	k8s.io/kubectl v0.27.3
)

require (
	github.com/btwiuse/gods v0.0.1 // indirect
	github.com/btwiuse/pretty v0.2.1 // indirect
	github.com/btwiuse/sse v0.0.1 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/dennwc/btrfs v0.0.0-20221026161108-3097362dc072 // indirect
	github.com/dennwc/ioctl v1.0.0 // indirect
	github.com/ebi-yade/altsvc-go v0.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-gost/gosocks4 v0.0.1 // indirect
	github.com/go-gost/gosocks5 v0.3.0 // indirect
	github.com/go-gost/relay v0.1.1-0.20211123134818-8ef7fd81ffd7 // indirect
	github.com/go-gost/tls-dissector v0.0.2-0.20220408131628-aac992c27451 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/pprof v0.0.0-20230821062121-407c9e7a662f // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.0.0-20220520183353-fd19c99a87aa // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/kataras/basicauth v0.0.3 // indirect
	github.com/mdlayher/ethtool v0.0.0-20220830195143-0e16326d06d1 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/monochromegane/go-gitignore v0.0.0-20200626010858-205db1a8cc00 // indirect
	github.com/onsi/ginkgo/v2 v2.12.0 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/quic-go v0.41.0 // indirect
	github.com/quic-go/webtransport-go v0.6.0 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tklauser/go-sysconf v0.3.5 // indirect
	github.com/tklauser/numcpus v0.2.2 // indirect
	github.com/webteleport/auth v0.0.6 // indirect
	github.com/webteleport/utils v0.2.2 // indirect
	github.com/webteleport/webteleport v0.3.3 // indirect
	github.com/webteleport/wtf v0.0.6 // indirect
	github.com/xlab/treeprint v1.1.0 // indirect
	gitlab.com/yawning/edwards25519-extra.git v0.0.0-20211229043746-2f91fcc9fbdb // indirect
	go.etcd.io/etcd/api/v3 v3.5.7 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.7 // indirect
	go.uber.org/mock v0.3.0 // indirect
	golang.org/x/exp v0.0.0-20240119083558-1b970713d09a // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	k0s.io/pkg/dial v0.0.0-00010101000000-000000000000 // indirect
	k0s.io/pkg/ui v0.1.5 // indirect
	k8s.io/component-base v0.27.3 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/kustomize/api v0.13.2 // indirect
	sigs.k8s.io/kustomize/kustomize/v5 v5.0.1 // indirect
	sigs.k8s.io/kustomize/kyaml v0.14.1 // indirect
)

require (
	cloud.google.com/go/compute v1.7.0 // indirect
	filippo.io/edwards25519 v1.0.0-rc.1.0.20210721174708-390f27c3be20 // indirect
	git.torproject.org/pluggable-transports/goptlib.git v1.2.0 // indirect
	gitea.com/lunny/log v0.0.0-20190322053110-01b5df579c4e // indirect
	github.com/ActiveState/termtest/conpty v0.5.0 // indirect
	github.com/AndreasBriese/bbloom v0.0.0-20190825152654-46b345b51c96 // indirect
	github.com/Azure/azure-sdk-for-go v58.0.0+incompatible // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.27 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.20 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.8 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.2 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/Azure/go-ntlmssp v0.0.0-20200615164410-66371956d46c // indirect
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/DataDog/datadog-go v3.7.2+incompatible // indirect
	github.com/GeertJohan/go.rice v1.0.0 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/LiamHaworth/go-tproxy v0.0.0-20190726054950-ef7efd7f24ed // indirect
	github.com/MakeNowJust/heredoc v1.0.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/Masterminds/sprig/v3 v3.2.2 // indirect
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/Rican7/retry v0.1.0 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/VojtechVitek/yaml-cli v0.0.5 // indirect
	github.com/abbot/go-http-auth v0.4.0 // indirect
	github.com/abiosoft/ishell v2.0.0+incompatible // indirect
	github.com/abiosoft/readline v0.0.0-20180607040430-155bce2042db // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/alecthomas/chroma v0.10.0 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/andrew-d/go-termutil v0.0.0-20150726205930-009166a695a2 // indirect
	github.com/armon/go-metrics v0.3.9 // indirect
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5 // indirect
	github.com/aryann/difflib v0.0.0-20210328193216-ff5ff6dc229b // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/asdine/storm v2.1.2+incompatible // indirect
	github.com/aws/aws-sdk-go v1.41.14 // indirect
	github.com/beevik/ntp v0.3.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/buildkite/interpolate v0.0.0-20200526001904-07f35b4ae251 // indirect
	github.com/buildkite/shellwords v0.0.0-20180315084142-c3f497d1e000 // indirect
	github.com/buildkite/yaml v0.0.0-20181016232759-0caa5f0796e3 // indirect
	github.com/caddyserver/caddy v1.0.5 // indirect
	github.com/caddyserver/certmagic v0.16.1 // indirect
	github.com/casbin/casbin/v2 v2.8.6 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/chai2010/gettext-go v1.0.2 // indirect
	github.com/checkpoint-restore/go-criu/v5 v5.3.0 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/cilium/ebpf v0.9.3 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/containerd/containerd v1.6.1 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.4.1 // indirect
	github.com/containerd/ttrpc v1.1.0 // indirect
	github.com/coreos/go-iptables v0.6.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cppforlife/cobrautil v0.0.0-20200514214827-bb86e6965d72 // indirect
	github.com/cppforlife/color v1.9.1-0.20200716202919-6706ac40b835 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/cyphar/filepath-securejoin v0.2.3 // indirect
	github.com/daaku/go.zipexe v1.0.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/daviddengcn/go-colortext v1.0.0 // indirect
	github.com/dchest/siphash v1.2.2 // indirect
	github.com/denisbrodbeck/machineid v1.0.1 // indirect
	github.com/dgraph-io/badger v1.6.2 // indirect
	github.com/dgraph-io/badger/v2 v2.2007.4 // indirect
	github.com/dgraph-io/ristretto v0.0.4-0.20200906165740-41ebdbffecfd // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dgryski/go-metro v0.0.0-20200812162917-85c65e2d0165 // indirect
	github.com/digitalocean/godo v1.41.0 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dlclark/regexp2 v1.4.1-0.20201116162257-a2a8dda75c91 // indirect
	github.com/docker/cli v20.10.9+incompatible // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/docker v20.10.17+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/dustin/go-humanize v1.0.1-0.20200219035652-afde56e7acac // indirect
	github.com/elithrar/simple-scrypt v1.3.0 // indirect
	github.com/ema/qdisc v0.0.0-20200603082823-62d0308e3e00 // indirect
	github.com/euank/go-kmsg-parser v2.0.0+incompatible // indirect
	github.com/evanphx/json-patch v4.12.0+incompatible // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/flynn/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/fvbommel/sortorder v1.0.1 // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell v1.4.0 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/go-acme/lego/v3 v3.7.0 // indirect
	github.com/go-asn1-ber/asn1-ber v1.5.1 // indirect
	github.com/go-chi/chi v4.1.2+incompatible // indirect
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/go-chi/cors v1.2.1 // indirect
	github.com/go-chi/render v1.0.1 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-ldap/ldap/v3 v3.4.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.1 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/cel-go v0.12.6 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-containerregistry v0.5.1 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/gax-go/v2 v2.4.0 // indirect
	github.com/goproxyio/windows v0.0.0-20191126033816-f4a809841617 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/hashicorp/cronexpr v1.1.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-envparse v0.1.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-syslog v1.0.0 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/memberlist v0.1.4 // indirect
	github.com/hashicorp/nomad/api v0.0.0-20220211135303-4afc67b7002e // indirect
	github.com/hashicorp/serf v0.8.3 // indirect
	github.com/hodgesds/perf-utils v0.5.1 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/illumos/go-kstat v0.0.0-20210513183136-173c9b0a9973 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/infobloxopen/go-trees v0.0.0-20200715205103-96a057b8dfb9 // indirect
	github.com/insomniacslk/dhcp v0.0.0-20220504074936-1ca156eafb9f // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.9.0 // indirect
	github.com/jackc/pgx/v4 v4.14.0 // indirect
	github.com/jaypipes/ghw v0.8.1-0.20220131141055-fb0598ce62c8 // indirect
	github.com/jaypipes/pcidb v0.6.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/josharian/native v1.0.0 // indirect
	github.com/jpillora/ansi v1.0.2 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/jpillora/chisel v1.7.7 // indirect
	github.com/jpillora/cookieauth v1.0.0 // indirect
	github.com/jpillora/eventsource v1.0.0 // indirect
	github.com/jpillora/go-echo-server v0.5.0 // indirect
	github.com/jpillora/ipfilter v1.2.1 // indirect
	github.com/jpillora/requestlog v1.0.0 // indirect
	github.com/jpillora/sizestr v1.0.0 // indirect
	github.com/jpillora/velox v0.4.0 // indirect
	github.com/jsimonetti/rtnetlink v1.3.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/klauspost/compress v1.15.15 // indirect
	github.com/klauspost/cpuid v1.3.1 // indirect
	github.com/klauspost/cpuid/v2 v2.1.1 // indirect
	github.com/klauspost/reedsolomon v1.9.15 // indirect
	github.com/koding/websocketproxy v0.0.0-20181220232114-7ed82d81a28c // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/libdns/alidns v1.0.2-x2 // indirect
	github.com/libdns/azure v0.2.0 // indirect
	github.com/libdns/cloudflare v0.1.0 // indirect
	github.com/libdns/digitalocean v0.0.0-20210310230526-186c4ebd2215 // indirect
	github.com/libdns/dnspod v0.0.3 // indirect
	github.com/libdns/duckdns v0.1.1 // indirect
	github.com/libdns/libdns v0.2.1 // indirect
	github.com/libdns/route53 v1.1.2 // indirect
	github.com/libdns/vultr v0.0.0-20211122184636-cd4cb5c12e51 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/lucas-clemente/quic-go v0.28.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.0.3 // indirect
	github.com/lufia/iostat v1.2.1 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/manifoldco/promptui v0.9.0 // indirect
	github.com/marten-seemann/qpack v0.2.1 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.5 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-18 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-19 v0.1.0 // indirect
	github.com/maruel/natural v0.0.0-20180416170133-dbcb3e2e8cf1 // indirect
	github.com/marusama/semaphore/v2 v2.4.1 // indirect
	github.com/mastercactapus/proxyprotocol v0.0.3 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-shellwords v1.0.12 // indirect
	github.com/mattn/go-xmlrpc v0.0.3 // indirect
	github.com/mattn/go-zglob v0.0.1 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2 // indirect
	github.com/mdlayher/genetlink v1.2.0 // indirect
	github.com/mdlayher/netlink v1.7.0 // indirect
	github.com/mdlayher/socket v0.4.0 // indirect
	github.com/mdlayher/wifi v0.0.0-20220330172155-a44c70b6d3c8 // indirect
	github.com/mesos/mesos-go v0.0.11 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mholt/acmez v1.0.2 // indirect
	github.com/mholt/archiver v3.1.1+incompatible // indirect
	github.com/mholt/certmagic v0.8.3 // indirect
	github.com/micromdm/scep/v2 v2.1.0 // indirect
	github.com/miekg/dns v1.1.58 // indirect
	github.com/mindprince/gonvml v0.0.0-20190828220739-9ebdce4bb989 // indirect
	github.com/mistifyio/go-zfs v2.1.2-0.20190413222219-f784269be439+incompatible // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/sys/mountinfo v0.5.0 // indirect
	github.com/moby/term v0.0.0-20221205130635-1aeaba878587 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mrunalp/fileutils v0.5.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416 // indirect
	github.com/nightlyone/lockfile v0.0.0-20180618180623-0ad87eef1443 // indirect
	github.com/nrdcg/dnspod-go v0.4.0 // indirect
	github.com/nwaples/rardecode v1.0.0 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oleiade/reflections v0.0.0-20160817071559-0e86b3c98b2f // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runc v1.1.0 // indirect
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417 // indirect
	github.com/opencontainers/selinux v1.10.2 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/oschwald/geoip2-golang v1.7.0 // indirect
	github.com/oschwald/maxminddb-golang v1.9.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/phuslu/geoip v1.0.20200217 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pires/go-proxyproto v0.5.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/portainer/docker-compose-wrapper v0.0.0-20220708023447-a69a4ebaa021 // indirect
	github.com/portainer/libcrypto v0.0.0-20210422035235-c652195c5c3a // indirect
	github.com/portainer/libhttp v0.0.0-20211208103139-07a5f798eb3f // indirect
	github.com/portainer/portainer/api v0.0.0-20220303203420-547d9c2fde15 // indirect
	github.com/posener/complete v1.2.2-0.20190308074557-af07aa5181b3 // indirect
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/node_exporter v1.5.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/qri-io/jsonpointer v0.0.0-20180309164927-168dd9e45cf2 // indirect
	github.com/qri-io/jsonschema v0.0.0-20180607150648-d0d3b10ec792 // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/refraction-networking/utls v0.0.0-20200601200209-ada0bb9b38a0 // indirect
	github.com/riobard/go-bloom v0.0.0-20200614022211-cdc8013cb5b3 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/rs/xid v1.2.1 // indirect
	github.com/russross/blackfriday v1.5.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/safchain/ethtool v0.2.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
	github.com/seccomp/libseccomp-golang v0.9.2-0.20210429002308-3879420cc921 // indirect
	github.com/seiflotfy/cuckoofilter v0.0.0-20201222105146-bc6005554a0c // indirect
	github.com/shadowsocks/go-shadowsocks2 v0.1.5 // indirect
	github.com/shadowsocks/shadowsocks-go v0.0.0-20200409064450-3e585ff90601 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/slackhq/nebula v1.5.2 // indirect
	github.com/smallstep/certificates v0.19.0 // indirect
	github.com/smallstep/cli v0.18.0 // indirect
	github.com/smallstep/nosql v0.4.0 // indirect
	github.com/smallstep/truststore v0.11.0 // indirect
	github.com/songgao/water v0.0.0-20200317203138-2b4b6d7c09d8 // indirect
	github.com/soundcloud/go-runit v0.0.0-20150630195641-06ad41a06c4a // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/cobra v1.6.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.0 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/tailscale/tscert v0.0.0-20220316030059-54bbcb9f74e2 // indirect
	github.com/templexxx/cpu v0.0.7 // indirect
	github.com/templexxx/xorsimd v0.4.1 // indirect
	github.com/tidwall/gjson v1.14.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tinylib/msgp v1.1.2 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce // indirect
	github.com/txthinking/runnergroup v0.0.0-20200327135940-540a793bb997 // indirect
	github.com/txthinking/socks5 v0.0.0-20200531111549-252709fcb919 // indirect
	github.com/txthinking/x v0.0.0-20200330144832-5ad2416896a9 // indirect
	github.com/u-root/uio v0.0.0-20210528114334-82958018845c // indirect
	github.com/ulikunitz/xz v0.5.8 // indirect
	github.com/vishvananda/netlink v1.1.1-0.20210330154013-f5de75959ad5 // indirect
	github.com/vishvananda/netns v0.0.0-20211101163701-50045581ed74 // indirect
	github.com/vito/go-interact v0.0.0-20171111012221-fa338ed9e9ec // indirect
	github.com/vultr/govultr/v2 v2.11.0 // indirect
	github.com/wI2L/jsondiff v0.2.0 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	github.com/xiaq/persistent v0.0.0-20200820214153-3175cfb92e14 // indirect
	github.com/xtaci/kcp-go/v5 v5.6.1 // indirect
	github.com/xtaci/smux v1.5.17 // indirect
	github.com/xtaci/tcpraw v1.2.25 // indirect
	github.com/yuin/goldmark v1.4.13 // indirect
	github.com/yuin/goldmark-highlighting v0.0.0-20220208100518-594be1970594 // indirect
	gitlab.com/mjwhitta/errors v1.0.0 // indirect
	gitlab.com/mjwhitta/hilighter v1.11.0 // indirect
	gitlab.com/mjwhitta/pathname v1.2.0 // indirect
	gitlab.com/mjwhitta/safety v1.11.0 // indirect
	gitlab.com/mjwhitta/sysinfo v1.4.7 // indirect
	gitlab.com/mjwhitta/where v1.2.4 // indirect
	gitlab.com/yawning/obfs4.git v0.0.0-20220904064028-336a71d6e4cf // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.mozilla.org/pkcs7 v0.0.0-20210826202110-33d05740a352 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.35.1 // indirect
	go.opentelemetry.io/otel v1.10.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.10.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.10.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.10.0 // indirect
	go.opentelemetry.io/otel/metric v0.31.0 // indirect
	go.opentelemetry.io/otel/sdk v1.10.0 // indirect
	go.opentelemetry.io/otel/trace v1.10.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.step.sm/cli-utils v0.7.0 // indirect
	go.step.sm/crypto v0.16.1 // indirect
	go.step.sm/linkedca v0.15.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/oauth2 v0.0.0-20220909003341-f21342109be1 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/term v0.16.0 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	golang.org/x/tools v0.17.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	gomodules.xyz/jsonpatch/v3 v3.0.1 // indirect
	gomodules.xyz/orderedmap v0.1.0 // indirect
	google.golang.org/api v0.84.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220706185917-7780775163c4 // indirect
	google.golang.org/grpc v1.51.0 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.28.0 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.56.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	howett.net/plist v1.0.0 // indirect
	// k0s.io/pkg/api v0.0.0-20220709103107-f40561c09b04 // indirect
	k0s.io/pkg/client v0.1.5 // indirect
	k0s.io/pkg/distro v0.1.5 // indirect
	k0s.io/pkg/fzf v0.1.5 // indirect
	k0s.io/pkg/hub v0.1.5 // indirect
	k0s.io/pkg/jsondns v0.1.5 // indirect
	k0s.io/pkg/manager v0.1.5 // indirect
	k0s.io/pkg/middleware v0.1.6 // indirect
	k0s.io/pkg/simple v0.1.5 // indirect
	// k0s.io/pkg/version v0.0.0-20220709103107-f40561c09b04 // indirect
	k0s.io/pkg/wrap v0.1.6 // indirect
	k8s.io/cli-runtime v0.27.3 // indirect
	k8s.io/component-helpers v0.27.3 // indirect
	k8s.io/kube-openapi v0.0.0-20230501164219-8b0f38b5fd1f // indirect
	k8s.io/metrics v0.27.3 // indirect
	k8s.io/utils v0.0.0-20230209194617-a36077c30491 // indirect
	nhooyr.io/websocket v1.8.10 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
	v2ray.com/core v4.19.1+incompatible // indirect
)

replace k0s.io => ../

replace k0s.io/cmd => ../cmd/

replace k0s.io/pkg/agent => ../pkg/agent/

replace k0s.io/pkg/asciitransport => ../pkg/asciitransport/

replace k0s.io/pkg/cli => ../pkg/cli/

replace k0s.io/pkg/client => ../pkg/client/

replace k0s.io/pkg/console => ../pkg/console/

replace k0s.io/pkg/distro => ../pkg/distro/

replace k0s.io/pkg/dohserver => ../pkg/dohserver/

replace k0s.io/pkg/exporter => ../pkg/exporter/

replace k0s.io/pkg/fonts => ../pkg/fonts/

replace k0s.io/pkg/fzf => ../pkg/fzf/

replace k0s.io/pkg/gitd => ../pkg/gitd/

replace k0s.io/pkg/hub => ../pkg/hub/

replace k0s.io/pkg/jsondns => ../pkg/jsondns/

replace k0s.io/pkg/manager => ../pkg/manager/

replace k0s.io/pkg/middleware => ../pkg/middleware/

replace k0s.io/pkg/plugin => ../pkg/plugin/

replace k0s.io/pkg/reverseproxy => ../pkg/reverseproxy/

replace k0s.io/pkg/rng => ../pkg/rng/

replace k0s.io/pkg/simple => ../pkg/simple/

replace k0s.io/pkg/tunnel => ../pkg/tunnel/

replace k0s.io/pkg/ui => ../pkg/ui/

replace k0s.io/pkg/utils => ../pkg/utils/

replace k0s.io/pkg/uuid => ../pkg/uuid/

replace k0s.io/pkg/wrap => ../pkg/wrap/

replace k0s.io/third_party => ./

replace k0s.io/pkg/dial => ../pkg/dial/
