module k0s.io

go 1.16

replace (
	github.com/buildkite/agent/v3 => github.com/btwiuse/agent/v3 v3.27.1-0.20210217080418-ae42a28eefa7
	github.com/caddyserver/forwardproxy => github.com/klzgrad/forwardproxy v0.0.0-20210120121422-9b4a5a242bd6
	github.com/ginuerzh/gost => github.com/btwiuse/gost v0.0.0-20210218034515-4e5ef1691e9f
	github.com/greenpau/caddy-auth-portal => github.com/btwiuse/caddy-auth-portal v1.3.12-0.20210204101408-068c2618b417
	k8s.io/apiserver => github.com/btwiuse/apiserver v0.0.0-20210224090623-3008f508f679
	k8s.io/kube-state-metrics/v2 => github.com/btwiuse/k16s/v2 v2.0.0-beta.0.20201224174453-2114e07844a9
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
	github.com/btwiuse/gods v0.0.0-20190414062120-7e7cf0aebbb0
	github.com/btwiuse/pretty v0.0.0-20190401073227-519ff4ea1882
	github.com/btwiuse/wetty v0.0.36
	github.com/buildkite/agent/v3 v3.27.0
	github.com/caddyserver/caddy/v2 v2.3.0
	github.com/caddyserver/forwardproxy v0.0.0-00010101000000-000000000000
	github.com/caddyserver/nginx-adapter v0.0.3
	github.com/casbin/caddy-authz/v2 v2.0.0
	github.com/creack/pty v1.1.11
	github.com/dchest/siphash v1.2.1 // indirect
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/docker/docker v1.4.2-0.20200214221943-d8772509d1a2
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/filebrowser/filebrowser/v2 v2.11.0
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/freman/caddy2-reauth v0.0.0-20200518130136-6064aa96b1a8
	github.com/gdamore/tcell v1.3.0
	github.com/ginuerzh/gost v0.0.0-20210206051340-8dd4d8d9a123
	github.com/go-log/log v0.2.0
	github.com/golang/protobuf v1.4.3
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
	github.com/jpillora/opts v1.2.0
	github.com/jpillora/webproc v0.4.0
	github.com/json-iterator/go v1.1.10
	github.com/klauspost/cpuid v1.3.1 // indirect
	github.com/klauspost/cpuid/v2 v2.0.3 // indirect
	github.com/liamg/aminal v0.9.0
	github.com/lucasb-eyer/go-colorful v1.0.3 // indirect
	github.com/lukesampson/figlet v0.0.0-20190211215653-8a3ef4a6ac42
	github.com/lunny/tango v0.5.6
	github.com/mattn/go-isatty v0.0.12
	github.com/mattn/go-runewidth v0.0.9
	github.com/mattn/go-shellwords v1.0.10
	github.com/mholt/caddy-l4 v0.0.0-20210209073014-d1d54b015e34
	github.com/mholt/caddy-webdav v0.0.0-20200916200058-c949b3226234
	github.com/miekg/dns v1.1.35
	github.com/p4gefau1t/trojan-go v0.8.2
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0
	github.com/prometheus/common v0.15.0
	github.com/prometheus/node_exporter v1.0.1
	github.com/pupapaik/sysinfo v0.0.0-20200106202926-c17dea004cd4
	github.com/rancher/dapper v0.5.5
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.8.0
	github.com/stretchr/testify v1.7.0
	github.com/tango-contrib/basicauth v0.0.0-20170526072948-7fbc19aece86
	github.com/traefik/yaegi v0.9.14
	github.com/urfave/cli v1.22.5
	go.starlark.net v0.0.0-20210212215732-ebe61bd709bf
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/mod v0.4.1
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	golang.org/x/sys v0.0.0-20210218085108-9555bcde0c6a
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/apiserver v0.19.7
	k8s.io/autoscaler/vertical-pod-autoscaler v0.9.2
	k8s.io/client-go v0.20.4
	k8s.io/component-base v0.19.7
	k8s.io/klog/v2 v2.5.0
	k8s.io/kube-state-metrics/v2 v2.0.0-00010101000000-000000000000
	modernc.org/httpfs v1.0.0
	nhooyr.io/websocket v1.8.6
	robpike.io/ivy v0.1.0
	sigs.k8s.io/metrics-server v0.4.2
	src.elv.sh v0.14.1-0.20210218105754-53593c3ab79f
)
