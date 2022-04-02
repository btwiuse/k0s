module k0s.io

go 1.17

replace (
	github.com/buildkite/agent/v3 => github.com/btwiuse/agent/v3 v3.27.1-0.20210217080418-ae42a28eefa7
	github.com/caddyserver/forwardproxy => github.com/klzgrad/forwardproxy v0.0.0-20210120121422-9b4a5a242bd6
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
	github.com/alexpantyukhin/go-pattern-match v0.0.0-20200628201436-c57d5ad3f2c5
	github.com/btwiuse/bcrypt v1.0.2
	github.com/btwiuse/cadvisor v0.0.0-20210312172035-34fddda41018
	github.com/btwiuse/dkg v0.2.0
	github.com/btwiuse/etcd/v3 v3.4.15
	github.com/btwiuse/gods v0.0.0-20190414062120-7e7cf0aebbb0
	github.com/btwiuse/pretty v0.0.0-20190401073227-519ff4ea1882
	github.com/btwiuse/wetty v0.0.36
	github.com/buildkite/agent/v3 v3.27.0
	github.com/caddy-dns/alidns v1.0.3-0.20210322055141-7532f6617482
	github.com/caddy-dns/azure v0.2.0
	github.com/caddy-dns/cloudflare v0.0.0-20210105070211-eda8e5aa2223
	github.com/caddy-dns/digitalocean v0.0.0-20210408173619-385f9346b5ac
	github.com/caddy-dns/dnspod v0.0.1
	github.com/caddy-dns/duckdns v0.2.2
	github.com/caddy-dns/route53 v1.1.1
	github.com/caddy-dns/vultr v0.0.0-20210105121117-3162aa6b9c27
	github.com/caddyserver/caddy/v2 v2.4.6
	github.com/caddyserver/forwardproxy v0.0.0-00010101000000-000000000000
	github.com/casbin/caddy-authz/v2 v2.0.0
	github.com/coredns/coredns v1.8.4-0.20210224180316-9d3a84377cae
	github.com/cppforlife/go-cli-ui v0.0.0-20200716203538-1e47f820817f
	github.com/creack/pty v1.1.11
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/docker/docker v20.10.2+incompatible
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/filebrowser/filebrowser/v2 v2.11.0
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/freman/caddy2-reauth v0.0.0-20200518130136-6064aa96b1a8
	github.com/gdamore/tcell v1.3.0
	github.com/ginuerzh/gost v0.0.0-20220402134819-e159aace1eaa
	github.com/go-log/log v0.2.0
	github.com/golang/protobuf v1.5.2
	github.com/google/cadvisor v0.39.0
	github.com/google/uuid v1.3.0
	github.com/goproxyio/goproxy/v2 v2.0.5
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/greenpau/caddy-auth-jwt v1.2.6
	github.com/greenpau/caddy-auth-portal v0.0.0-00010101000000-000000000000
	github.com/greenpau/caddy-trace v1.1.5
	github.com/hairyhenderson/caddy-teapot-module v0.0.2
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/iamd3vil/caddy_yaml_adapter v0.0.0-20200503183711-d479c29b475a
	github.com/infobloxopen/go-trees v0.0.0-20200715205103-96a057b8dfb9
	github.com/jaypipes/ghw v0.7.0 // indirect
	github.com/jpillora/chisel v1.7.6 // indirect
	github.com/jpillora/opts v1.2.0
	github.com/jpillora/webproc v0.4.0
	github.com/json-iterator/go v1.1.11
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/koding/websocketproxy v0.0.0-20181220232114-7ed82d81a28c // indirect
	github.com/liamg/aminal v0.9.0
	github.com/lukesampson/figlet v0.0.0-20190211215653-8a3ef4a6ac42
	github.com/lunny/tango v0.5.6
	github.com/mathetake/gasm v0.0.0-20200928142744-80e74517647c
	github.com/mattn/go-isatty v0.0.13
	github.com/mattn/go-runewidth v0.0.10
	github.com/mattn/go-shellwords v1.0.10
	github.com/mholt/caddy-l4 v0.0.0-20210209073014-d1d54b015e34
	github.com/mholt/caddy-webdav v0.0.0-20200916200058-c949b3226234
	github.com/miekg/dns v1.1.43
	github.com/p4gefau1t/trojan-go v0.8.2
	github.com/pkg/errors v0.9.1
	github.com/portainer/agent v0.0.0-20210129020346-65fa249df79a
	github.com/portainer/libcrypto v0.0.0-20201216185936-7a703a1ea452 // indirect
	github.com/portainer/libhttp v0.0.0-20201216185909-d20481a3da82 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.26.0
	github.com/prometheus/node_exporter v1.0.1
	github.com/rancher/dapper v0.5.5
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.8.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tango-contrib/basicauth v0.0.0-20170526072948-7fbc19aece86
	github.com/traefik/yaegi v0.9.14
	github.com/urfave/cli v1.22.5
	github.com/v2fly/v2ray-core/v4 v4.36.2
	gitlab.com/mjwhitta/sysinfo v1.2.5
	go.starlark.net v0.0.0-20210305151048-6a590ae7f4eb
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3
	golang.org/x/mod v0.5.1
	golang.org/x/net v0.0.0-20211215060638-4ddde0e984e9
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e
	golang.org/x/text v0.3.7
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
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
	modernc.org/httpfs v1.0.0
	nhooyr.io/websocket v1.8.6
	robpike.io/ivy v0.1.0
	sigs.k8s.io/metrics-server v0.4.2
	src.elv.sh v0.14.1-0.20210218105754-53593c3ab79f
)

require (
	cloud.google.com/go v0.83.0 // indirect
	git.torproject.org/pluggable-transports/goptlib.git v1.2.0 // indirect
	gitea.com/lunny/log v0.0.0-20190322053110-01b5df579c4e // indirect
	github.com/AndreasBriese/bbloom v0.0.0-20190825152654-46b345b51c96 // indirect
	github.com/Azure/azure-sdk-for-go v52.4.0+incompatible // indirect
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
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
	github.com/GeertJohan/go.rice v1.0.0 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/LiamHaworth/go-tproxy v0.0.0-20190726054950-ef7efd7f24ed // indirect
	github.com/MakeNowJust/heredoc v0.0.0-20170808103936-bb23615498cd // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/Masterminds/sprig/v3 v3.2.2 // indirect
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/Rican7/retry v0.1.0 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/aead/chacha20 v0.0.0-20180709150244-8b13a72661da // indirect
	github.com/alecthomas/chroma v0.9.2 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/andrew-d/go-termutil v0.0.0-20150726205930-009166a695a2 // indirect
	github.com/antlr/antlr4 v0.0.0-20200503195918-621b933c7a7f // indirect
	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/asdine/storm v2.1.2+incompatible // indirect
	github.com/aws/aws-sdk-go v1.37.10 // indirect
	github.com/beevik/etree v1.1.0 // indirect
	github.com/beevik/ntp v0.3.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/buildkite/interpolate v0.0.0-20200526001904-07f35b4ae251 // indirect
	github.com/buildkite/shellwords v0.0.0-20180315084142-c3f497d1e000 // indirect
	github.com/buildkite/yaml v0.0.0-20181016232759-0caa5f0796e3 // indirect
	github.com/caddyserver/caddy v1.0.3 // indirect
	github.com/caddyserver/certmagic v0.15.2 // indirect
	github.com/casbin/casbin/v2 v2.8.6 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/chai2010/gettext-go v0.0.0-20160711120539-c6fed771bfd5 // indirect
	github.com/checkpoint-restore/go-criu/v4 v4.1.0 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/cilium/ebpf v0.2.0 // indirect
	github.com/containerd/console v1.0.1 // indirect
	github.com/containerd/containerd v1.4.4 // indirect
	github.com/containerd/ttrpc v1.0.2 // indirect
	github.com/coredns/caddy v1.1.0 // indirect
	github.com/coreos/go-iptables v0.6.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/cppforlife/cobrautil v0.0.0-20200514214827-bb86e6965d72 // indirect
	github.com/cppforlife/color v1.9.1-0.20200716202919-6706ac40b835 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/crewjam/httperr v0.0.0-20190612203328-a946449404da // indirect
	github.com/crewjam/saml v0.4.5 // indirect
	github.com/cyphar/filepath-securejoin v0.2.2 // indirect
	github.com/daaku/go.zipexe v1.0.1 // indirect
	github.com/danwakefield/fnmatch v0.0.0-20160403171240-cbb64ac3d964 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/daviddengcn/go-colortext v0.0.0-20160507010035-511bcaf42ccd // indirect
	github.com/dchest/siphash v1.2.2 // indirect
	github.com/dgraph-io/badger v1.6.2 // indirect
	github.com/dgraph-io/badger/v2 v2.2007.4 // indirect
	github.com/dgraph-io/ristretto v0.0.4-0.20200906165740-41ebdbffecfd // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dgryski/go-jump v0.0.0-20170409065014-e1f439676b57 // indirect
	github.com/dgryski/go-metro v0.0.0-20200812162917-85c65e2d0165 // indirect
	github.com/digitalocean/godo v1.41.0 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/dnstap/golang-dnstap v0.4.0 // indirect
	github.com/docker/cli v0.0.0-20191017083524-a8ff7f821017 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/docker/libcontainer v2.2.1+incompatible // indirect
	github.com/docker/spdystream v0.0.0-20160310174837-449fdfce4d96 // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/dustin/go-humanize v1.0.1-0.20200219035652-afde56e7acac // indirect
	github.com/elithrar/simple-scrypt v1.3.0 // indirect
	github.com/ema/qdisc v0.0.0-20190904071900-b82c76788043 // indirect
	github.com/emicklei/go-restful v2.15.0+incompatible // indirect
	github.com/euank/go-kmsg-parser v2.0.0+incompatible // indirect
	github.com/evanphx/json-patch v4.9.0+incompatible // indirect
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
	github.com/go-asn1-ber/asn1-ber v1.3.1 // indirect
	github.com/go-chi/chi v4.1.2+incompatible // indirect
	github.com/go-gost/gosocks4 v0.0.1 // indirect
	github.com/go-gost/gosocks5 v0.3.0 // indirect
	github.com/go-gost/relay v0.1.1-0.20211123134818-8ef7fd81ffd7 // indirect
	github.com/go-gost/tls-dissector v0.0.2-0.20211125135007-2b5d5bd9c07e // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-ldap/ldap v3.0.3+incompatible // indirect
	github.com/go-ldap/ldap/v3 v3.1.10 // indirect
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-logr/logr v0.4.0 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/godbus/dbus v0.0.0-20190402143921-271e53dc4968 // indirect
	github.com/godbus/dbus/v5 v5.0.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/cel-go v0.7.3 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/go-containerregistry v0.1.1 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/goproxyio/windows v0.0.0-20191126033816-f4a809841617 // indirect
	github.com/greenpau/go-identity v1.0.19 // indirect
	github.com/greenpau/versioned v1.0.23 // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.3 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.8 // indirect
	github.com/hashicorp/go-sockaddr v1.0.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/memberlist v0.1.3 // indirect
	github.com/hashicorp/serf v0.8.2 // indirect
	github.com/hodgesds/perf-utils v0.0.8 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
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
	github.com/juju/ansiterm v0.0.0-20180109212912-720a0952cc2a // indirect
	github.com/karrick/godirwalk v1.16.1 // indirect
	github.com/klauspost/cpuid v1.3.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/klauspost/reedsolomon v1.9.15 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/libdns/alidns v1.0.1 // indirect
	github.com/libdns/azure v0.2.0 // indirect
	github.com/libdns/cloudflare v0.0.0-20200528144945-97886e7873b1 // indirect
	github.com/libdns/digitalocean v0.0.0-20210310230526-186c4ebd2215 // indirect
	github.com/libdns/dnspod v0.0.1 // indirect
	github.com/libdns/duckdns v0.0.0-20201229185916-cd405ff644b9 // indirect
	github.com/libdns/libdns v0.2.1 // indirect
	github.com/libdns/route53 v1.1.0 // indirect
	github.com/libdns/vultr v0.0.0-20201128180404-1d5ee21ea62f // indirect
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/lucas-clemente/quic-go v0.26.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.0.2 // indirect
	github.com/lufia/iostat v1.1.0 // indirect
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/magefile/mage v1.10.0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/manifoldco/promptui v0.8.0 // indirect
	github.com/marten-seemann/qpack v0.2.1 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.5 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.1 // indirect
	github.com/marten-seemann/qtls-go1-18 v0.1.1 // indirect
	github.com/maruel/natural v0.0.0-20180416170133-dbcb3e2e8cf1 // indirect
	github.com/marusama/semaphore/v2 v2.4.1 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.0.0-20201213122252-bcd7e1b9601e // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-xmlrpc v0.0.3 // indirect
	github.com/mattn/go-zglob v0.0.1 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mdlayher/genetlink v1.0.0 // indirect
	github.com/mdlayher/netlink v1.1.0 // indirect
	github.com/mdlayher/wifi v0.0.0-20190303161829-b1436901ddee // indirect
	github.com/mesos/mesos-go v0.0.11 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	github.com/mholt/acmez v1.0.1 // indirect
	github.com/mholt/archiver v3.1.1+incompatible // indirect
	github.com/mholt/certmagic v0.6.2-0.20190624175158-6a42ef9fe8c2 // indirect
	github.com/micromdm/scep/v2 v2.1.0 // indirect
	github.com/milosgajdos/tenus v0.0.3 // indirect
	github.com/mindprince/gonvml v0.0.0-20190828220739-9ebdce4bb989 // indirect
	github.com/mistifyio/go-zfs v2.1.2-0.20190413222219-f784269be439+incompatible // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/sys/mountinfo v0.4.0 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/mrunalp/fileutils v0.5.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/naoina/toml v0.1.1 // indirect
	github.com/nightlyone/lockfile v0.0.0-20180618180623-0ad87eef1443 // indirect
	github.com/nrdcg/dnspod-go v0.4.0 // indirect
	github.com/nwaples/rardecode v1.0.0 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oleiade/reflections v0.0.0-20160817071559-0e86b3c98b2f // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v1.0.0-rc93 // indirect
	github.com/opencontainers/runtime-spec v1.0.3-0.20200929063507-e6143ca7d51d // indirect
	github.com/opencontainers/selinux v1.8.0 // indirect
	github.com/opentracing-contrib/go-observer v0.0.0-20170622124052-a52f23424492 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5 // indirect
	github.com/openzipkin/zipkin-go v0.2.2 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/phuslu/geoip v1.0.20200217 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pires/go-proxyproto v0.5.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/posener/complete v1.2.2-0.20190308074557-af07aa5181b3 // indirect
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/qri-io/jsonpointer v0.0.0-20180309164927-168dd9e45cf2 // indirect
	github.com/qri-io/jsonschema v0.0.0-20180607150648-d0d3b10ec792 // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/refraction-networking/utls v0.0.0-20200601200209-ada0bb9b38a0 // indirect
	github.com/riobard/go-bloom v0.0.0-20200614022211-cdc8013cb5b3 // indirect
	github.com/robfig/cron/v3 v3.0.0 // indirect
	github.com/rs/xid v1.2.1 // indirect
	github.com/russellhaering/goxmldsig v1.1.0 // indirect
	github.com/russross/blackfriday v1.5.2 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/samfoo/ansi v0.0.0-20160124022901-b6bd2ded7189 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
	github.com/seccomp/libseccomp-golang v0.9.1 // indirect
	github.com/seiflotfy/cuckoofilter v0.0.0-20201222105146-bc6005554a0c // indirect
	github.com/shadowsocks/go-shadowsocks2 v0.1.5 // indirect
	github.com/shadowsocks/shadowsocks-go v0.0.0-20200409064450-3e585ff90601 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/siebenmann/go-kstat v0.0.0-20200303194639-4e8294f9e9d5 // indirect
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e // indirect
	github.com/smallstep/certificates v0.17.5-0.20211008195551-04fe3126bebf // indirect
	github.com/smallstep/cli v0.17.6 // indirect
	github.com/smallstep/nosql v0.3.8 // indirect
	github.com/smallstep/truststore v0.9.6 // indirect
	github.com/songgao/water v0.0.0-20200317203138-2b4b6d7c09d8 // indirect
	github.com/soundcloud/go-runit v0.0.0-20150630195641-06ad41a06c4a // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/cobra v1.1.3 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.0 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/templexxx/xor v0.0.0-20191217153810-f85b25db303b // indirect
	github.com/tinylib/msgp v1.1.2 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce // indirect
	github.com/txthinking/runnergroup v0.0.0-20200327135940-540a793bb997 // indirect
	github.com/txthinking/socks5 v0.0.0-20200531111549-252709fcb919 // indirect
	github.com/txthinking/x v0.0.0-20200330144832-5ad2416896a9 // indirect
	github.com/ulikunitz/xz v0.5.7 // indirect
	github.com/vishvananda/netlink v1.1.0 // indirect
	github.com/vishvananda/netns v0.0.0-20200728191858-db3c7e526aae // indirect
	github.com/vito/go-interact v0.0.0-20171111012221-fa338ed9e9ec // indirect
	github.com/vultr/govultr/v2 v2.0.0 // indirect
	github.com/willf/bitset v1.1.11 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	github.com/xiaq/persistent v0.0.0-20200820214153-3175cfb92e14 // indirect
	github.com/xtaci/kcp-go v5.4.20+incompatible // indirect
	github.com/xtaci/smux v1.5.16 // indirect
	github.com/xtaci/tcpraw v1.2.25 // indirect
	github.com/yuin/goldmark v1.4.1 // indirect
	github.com/yuin/goldmark-highlighting v0.0.0-20210516132338-9216f9c5aa01 // indirect
	gitlab.com/mjwhitta/hilighter v1.9.2 // indirect
	gitlab.com/mjwhitta/pathname v1.0.9 // indirect
	gitlab.com/mjwhitta/safety v1.7.1 // indirect
	gitlab.com/mjwhitta/where v1.0.21 // indirect
	gitlab.com/yawning/obfs4.git v0.0.0-20210511220700-e330d1b7024b // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.mozilla.org/pkcs7 v0.0.0-20210826202110-33d05740a352 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.step.sm/cli-utils v0.6.0 // indirect
	go.step.sm/crypto v0.11.0 // indirect
	go.step.sm/linkedca v0.5.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.0 // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8 // indirect
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c // indirect
	golang.org/x/term v0.0.0-20210503060354-a79de5458b56 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	golang.org/x/tools v0.1.8 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gomodules.xyz/jsonpatch/v3 v3.0.1 // indirect
	gomodules.xyz/orderedmap v0.1.0 // indirect
	google.golang.org/api v0.48.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210719143636-1d5a45f8e492 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.28.0 // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.56.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	howett.net/plist v0.0.0-20200419221736-3b63eb3a43b5 // indirect
	k8s.io/cli-runtime v0.20.4 // indirect
	k8s.io/component-helpers v0.20.4 // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kube-openapi v0.0.0-20201113171705-d219536bb9fd // indirect
	k8s.io/metrics v0.20.4 // indirect
	k8s.io/utils v0.0.0-20210305010621-2afb4311ab10 // indirect
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.14 // indirect
	sigs.k8s.io/kustomize v2.0.3+incompatible // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.0.2 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
	v2ray.com/core v4.19.1+incompatible // indirect
)
