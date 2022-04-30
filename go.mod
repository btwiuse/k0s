module k0s.io

go 1.18

replace (
	github.com/buildkite/agent/v3 => github.com/btwiuse/agent/v3 v3.27.1-0.20210217080418-ae42a28eefa7
	github.com/coredns/coredns => github.com/btwiuse/coredns v1.8.4
	github.com/ginuerzh/gost => github.com/btwiuse/gost v0.0.0-20220402134819-e159aace1eaa
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
)

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/MaxRis/w32 v0.0.0-20180517000239-4f5cfb03fabf // indirect
	github.com/VojtechVitek/yaml-cli v0.0.5
	github.com/abbot/go-http-auth v0.4.0
	github.com/abiosoft/ishell v2.0.0+incompatible
	github.com/abiosoft/readline v0.0.0-20180607040430-155bce2042db // indirect
	github.com/alexpantyukhin/go-pattern-match v0.0.0-20200628201436-c57d5ad3f2c5
	github.com/btwiuse/bcrypt v1.0.2
	github.com/btwiuse/dkg v0.2.0
	github.com/btwiuse/etcd/v3 v3.4.15 // indirect
	github.com/btwiuse/gods v0.0.0-20190414062120-7e7cf0aebbb0
	github.com/btwiuse/pretty v0.0.0-20190401073227-519ff4ea1882
	github.com/btwiuse/wetty v0.0.36
	github.com/buildkite/agent/v3 v3.27.0
	github.com/caddyserver/caddy/v2 v2.4.6
	github.com/coredns/coredns v1.8.4-0.20210224180316-9d3a84377cae
	github.com/cppforlife/go-cli-ui v0.0.0-20200716203538-1e47f820817f
	github.com/creack/pty v1.1.11
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/docker/docker v20.10.9+incompatible
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/filebrowser/filebrowser/v2 v2.11.0
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/gdamore/tcell v1.3.0
	github.com/ginuerzh/gost v0.0.0-20220402134819-e159aace1eaa
	github.com/go-log/log v0.2.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/goproxyio/goproxy/v2 v2.0.5
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/infobloxopen/go-trees v0.0.0-20200715205103-96a057b8dfb9
	github.com/jaypipes/ghw v0.8.1-0.20220131141055-fb0598ce62c8 // indirect
	github.com/jpillora/chisel v1.7.7 // indirect
	github.com/jpillora/opts v1.2.0
	github.com/jpillora/webproc v0.4.0
	github.com/json-iterator/go v1.1.11
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/koding/websocketproxy v0.0.0-20181220232114-7ed82d81a28c // indirect
	github.com/liamg/aminal v0.9.0
	github.com/lukesampson/figlet v0.0.0-20190211215653-8a3ef4a6ac42
	github.com/lunny/tango v0.5.6
	github.com/mattn/go-isatty v0.0.13
	github.com/mattn/go-runewidth v0.0.10
	github.com/mattn/go-shellwords v1.0.10
	github.com/miekg/dns v1.1.47
	github.com/p4gefau1t/trojan-go v0.8.2
	github.com/pkg/errors v0.9.1
	github.com/portainer/agent v0.0.0-20220428235705-f483e4d7ecc4
	github.com/portainer/libcrypto v0.0.0-20210422035235-c652195c5c3a // indirect
	github.com/portainer/libhttp v0.0.0-20211208103139-07a5f798eb3f // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.26.0
	github.com/prometheus/node_exporter v1.0.1
	github.com/rancher/dapper v0.5.5
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.1
	github.com/tango-contrib/basicauth v0.0.0-20170526072948-7fbc19aece86
	github.com/traefik/yaegi v0.9.14
	github.com/urfave/cli v1.22.5
	github.com/v2fly/v2ray-core/v4 v4.36.2
	gitlab.com/mjwhitta/sysinfo v1.2.5
	go.starlark.net v0.0.0-20210305151048-6a590ae7f4eb
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd
	golang.org/x/mod v0.5.1
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20220315194320-039c03cc5b86
	golang.org/x/text v0.3.7
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	k8s.io/apiserver v0.20.6
	k8s.io/autoscaler/vertical-pod-autoscaler v0.9.2
	k8s.io/client-go v0.22.2
	k8s.io/component-base v0.20.6
	k8s.io/klog/v2 v2.60.1
	k8s.io/kube-state-metrics/v2 v2.4.2
	k8s.io/kubectl v0.20.4
	modernc.org/httpfs v1.0.0
	nhooyr.io/websocket v1.8.7
	robpike.io/ivy v0.1.0
	sigs.k8s.io/metrics-server v0.4.2
	src.elv.sh v0.14.1-0.20210218105754-53593c3ab79f
)

require (
	github.com/Dreamacro/clash v1.10.0
	go.uber.org/automaxprocs v1.5.1
)

require (
	cloud.google.com/go v0.83.0 // indirect
	git.torproject.org/pluggable-transports/goptlib.git v1.2.0 // indirect
	gitea.com/lunny/log v0.0.0-20190322053110-01b5df579c4e // indirect
	github.com/Azure/azure-sdk-for-go v52.4.0+incompatible // indirect
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.18 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.13 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.7 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.2 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/DataDog/datadog-go v3.7.2+incompatible // indirect
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/Dreamacro/go-shadowsocks2 v0.1.7 // indirect
	github.com/GeertJohan/go.rice v1.0.0 // indirect
	github.com/LiamHaworth/go-tproxy v0.0.0-20190726054950-ef7efd7f24ed // indirect
	github.com/MakeNowJust/heredoc v0.0.0-20170808103936-bb23615498cd // indirect
	github.com/Microsoft/go-winio v0.4.17 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15 // indirect
	github.com/andrew-d/go-termutil v0.0.0-20150726205930-009166a695a2 // indirect
	github.com/armon/go-metrics v0.0.0-20190430140413-ec5e00d3c878 // indirect
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/asdine/storm v2.1.2+incompatible // indirect
	github.com/aws/aws-sdk-go v1.37.10 // indirect
	github.com/beevik/ntp v0.3.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/buildkite/interpolate v0.0.0-20200526001904-07f35b4ae251 // indirect
	github.com/buildkite/shellwords v0.0.0-20180315084142-c3f497d1e000 // indirect
	github.com/buildkite/yaml v0.0.0-20181016232759-0caa5f0796e3 // indirect
	github.com/caddyserver/caddy v1.0.3 // indirect
	github.com/caddyserver/certmagic v0.15.2 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/chai2010/gettext-go v0.0.0-20160711120539-c6fed771bfd5 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/containerd/containerd v1.5.7 // indirect
	github.com/coredns/caddy v1.1.0 // indirect
	github.com/coreos/go-iptables v0.6.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/cppforlife/cobrautil v0.0.0-20200514214827-bb86e6965d72 // indirect
	github.com/cppforlife/color v1.9.1-0.20200716202919-6706ac40b835 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/daaku/go.zipexe v1.0.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/daviddengcn/go-colortext v0.0.0-20160507010035-511bcaf42ccd // indirect
	github.com/dchest/siphash v1.2.2 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-jump v0.0.0-20170409065014-e1f439676b57 // indirect
	github.com/dgryski/go-metro v0.0.0-20200812162917-85c65e2d0165 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dnstap/golang-dnstap v0.4.0 // indirect
	github.com/docker/cli v20.10.9+incompatible // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/docker/libcontainer v2.2.1+incompatible // indirect
	github.com/docker/spdystream v0.0.0-20160310174837-449fdfce4d96 // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/elithrar/simple-scrypt v1.3.0 // indirect
	github.com/ema/qdisc v0.0.0-20190904071900-b82c76788043 // indirect
	github.com/emicklei/go-restful v2.15.0+incompatible // indirect
	github.com/evanphx/json-patch v4.11.0+incompatible // indirect
	github.com/exponent-io/jsonpath v0.0.0-20151013193312-d6023ce2651d // indirect
	github.com/farsightsec/golang-framestream v0.3.0 // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/flynn/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/fvbommel/sortorder v1.0.1 // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/go-acme/lego v2.5.0+incompatible // indirect
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/go-chi/cors v1.2.0 // indirect
	github.com/go-chi/render v1.0.1 // indirect
	github.com/go-gost/gosocks4 v0.0.1 // indirect
	github.com/go-gost/gosocks5 v0.3.0 // indirect
	github.com/go-gost/relay v0.1.1-0.20211123134818-8ef7fd81ffd7 // indirect
	github.com/go-gost/tls-dissector v0.0.2-0.20211125135007-2b5d5bd9c07e // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/godbus/dbus v0.0.0-20190422162347-ade71ed3457e // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/go-containerregistry v0.1.1 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/goproxyio/windows v0.0.0-20191126033816-f4a809841617 // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645 // indirect
	github.com/hashicorp/cronexpr v1.1.1 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.1.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/memberlist v0.1.4 // indirect
	github.com/hashicorp/nomad/api v0.0.0-20220211135303-4afc67b7002e // indirect
	github.com/hashicorp/serf v0.8.3 // indirect
	github.com/hodgesds/perf-utils v0.0.8 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/insomniacslk/dhcp v0.0.0-20220119180841-3c283ff8b7dd // indirect
	github.com/jaypipes/pcidb v0.6.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/ansi v1.0.2 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/jpillora/cookieauth v1.0.0 // indirect
	github.com/jpillora/eventsource v1.0.0 // indirect
	github.com/jpillora/ipfilter v1.2.1 // indirect
	github.com/jpillora/requestlog v1.0.0 // indirect
	github.com/jpillora/sizestr v1.0.0 // indirect
	github.com/jpillora/velox v0.4.0 // indirect
	github.com/klauspost/cpuid v1.3.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/klauspost/reedsolomon v1.9.15 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/libdns/libdns v0.2.1 // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/lucas-clemente/quic-go v0.26.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.0.2 // indirect
	github.com/lufia/iostat v1.1.0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.5 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.1 // indirect
	github.com/marten-seemann/qtls-go1-18 v0.1.1 // indirect
	github.com/maruel/natural v0.0.0-20180416170133-dbcb3e2e8cf1 // indirect
	github.com/marusama/semaphore/v2 v2.4.1 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-xmlrpc v0.0.3 // indirect
	github.com/mattn/go-zglob v0.0.1 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mdlayher/genetlink v1.0.0 // indirect
	github.com/mdlayher/netlink v1.1.1 // indirect
	github.com/mdlayher/wifi v0.0.0-20190303161829-b1436901ddee // indirect
	github.com/mholt/acmez v1.0.1 // indirect
	github.com/mholt/archiver v3.1.1+incompatible // indirect
	github.com/mholt/certmagic v0.6.2-0.20190624175158-6a42ef9fe8c2 // indirect
	github.com/milosgajdos/tenus v0.0.3 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/naoina/toml v0.1.1 // indirect
	github.com/nightlyone/lockfile v0.0.0-20180618180623-0ad87eef1443 // indirect
	github.com/nwaples/rardecode v1.0.0 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oleiade/reflections v0.0.0-20160817071559-0e86b3c98b2f // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opentracing-contrib/go-observer v0.0.0-20170622124052-a52f23424492 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5 // indirect
	github.com/openzipkin/zipkin-go v0.2.2 // indirect
	github.com/oschwald/geoip2-golang v1.6.1 // indirect
	github.com/oschwald/maxminddb-golang v1.8.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/phuslu/geoip v1.0.20200217 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pires/go-proxyproto v0.5.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/portainer/docker-compose-wrapper v0.0.0-20220407011010-3c7408969ad3 // indirect
	github.com/portainer/portainer/api v0.0.0-20220303203420-547d9c2fde15 // indirect
	github.com/posener/complete v1.2.2-0.20190308074557-af07aa5181b3 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/qri-io/jsonpointer v0.0.0-20180309164927-168dd9e45cf2 // indirect
	github.com/qri-io/jsonschema v0.0.0-20180607150648-d0d3b10ec792 // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/refraction-networking/utls v0.0.0-20200601200209-ada0bb9b38a0 // indirect
	github.com/riobard/go-bloom v0.0.0-20200614022211-cdc8013cb5b3 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/russross/blackfriday v1.5.2 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
	github.com/seiflotfy/cuckoofilter v0.0.0-20201222105146-bc6005554a0c // indirect
	github.com/shadowsocks/go-shadowsocks2 v0.1.5 // indirect
	github.com/shadowsocks/shadowsocks-go v0.0.0-20200409064450-3e585ff90601 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/siebenmann/go-kstat v0.0.0-20200303194639-4e8294f9e9d5 // indirect
	github.com/songgao/water v0.0.0-20200317203138-2b4b6d7c09d8 // indirect
	github.com/soundcloud/go-runit v0.0.0-20150630195641-06ad41a06c4a // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/cobra v1.1.3 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20191217153810-f85b25db303b // indirect
	github.com/tinylib/msgp v1.1.2 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce // indirect
	github.com/txthinking/runnergroup v0.0.0-20200327135940-540a793bb997 // indirect
	github.com/txthinking/socks5 v0.0.0-20200531111549-252709fcb919 // indirect
	github.com/txthinking/x v0.0.0-20200330144832-5ad2416896a9 // indirect
	github.com/u-root/uio v0.0.0-20210528114334-82958018845c // indirect
	github.com/ulikunitz/xz v0.5.7 // indirect
	github.com/vito/go-interact v0.0.0-20171111012221-fa338ed9e9ec // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	github.com/xiaq/persistent v0.0.0-20200820214153-3175cfb92e14 // indirect
	github.com/xtaci/kcp-go v5.4.20+incompatible // indirect
	github.com/xtaci/smux v1.5.16 // indirect
	github.com/xtaci/tcpraw v1.2.25 // indirect
	gitlab.com/mjwhitta/hilighter v1.9.2 // indirect
	gitlab.com/mjwhitta/pathname v1.0.9 // indirect
	gitlab.com/mjwhitta/safety v1.7.1 // indirect
	gitlab.com/mjwhitta/where v1.0.21 // indirect
	gitlab.com/yawning/obfs4.git v0.0.0-20210511220700-e330d1b7024b // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.0 // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8 // indirect
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	golang.org/x/tools v0.1.8 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gomodules.xyz/jsonpatch/v3 v3.0.1 // indirect
	gomodules.xyz/orderedmap v0.1.0 // indirect
	google.golang.org/api v0.48.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210719143636-1d5a45f8e492 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.28.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.56.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	howett.net/plist v1.0.0 // indirect
	k8s.io/cli-runtime v0.20.4 // indirect
	k8s.io/component-helpers v0.20.4 // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kube-openapi v0.0.0-20210421082810-95288971da7e // indirect
	k8s.io/metrics v0.20.4 // indirect
	k8s.io/utils v0.0.0-20210819203725-bdf08cb9a70a // indirect
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.14 // indirect
	sigs.k8s.io/kustomize v2.0.3+incompatible // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.1.2 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
	v2ray.com/core v4.19.1+incompatible // indirect
)
