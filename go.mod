module k0s.io/k0s

go 1.15

// replace github.com/btwiuse/wetty => /home/aaron/wetty
// replace github.com/yrpc/yrpc => /home/aaron/yrpc/yrpc
replace nhooyr.io/websocket => github.com/btwiuse/websocket v1.8.6

replace github.com/milosgajdos83/tenus => github.com/milosgajdos/tenus v0.0.0-20200407174313-f075bef9ab28

require (
	github.com/MaxRis/w32 v0.0.0-20180517000239-4f5cfb03fabf // indirect
	github.com/VojtechVitek/yaml-cli v0.0.5
	github.com/abbot/go-http-auth v0.4.0
	github.com/abiosoft/ishell v2.0.0+incompatible
	github.com/abiosoft/readline v0.0.0-20180607040430-155bce2042db // indirect
	github.com/alexpantyukhin/go-pattern-match v0.0.0-20200628201436-c57d5ad3f2c5
	github.com/btwiuse/gods v0.0.0-20190414062120-7e7cf0aebbb0
	github.com/btwiuse/pretty v0.0.0-20190401073227-519ff4ea1882
	github.com/btwiuse/wetty v0.0.34
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/creack/pty v1.1.11
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/docker/docker v1.4.2-0.20200214221943-d8772509d1a2
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/flynn-archive/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/gdamore/tcell v1.3.0
	github.com/ginuerzh/gost v0.0.0-20200523132633-2707a8f0a90e
	github.com/go-log/log v0.2.0
	github.com/golang/protobuf v1.4.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/websocket v1.4.1
	github.com/json-iterator/go v1.1.9
	github.com/klauspost/reedsolomon v1.9.9 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/liamg/aminal v0.9.0
	github.com/lucasb-eyer/go-colorful v1.0.3 // indirect
	github.com/lukesampson/figlet v0.0.0-20190211215653-8a3ef4a6ac42
	github.com/mattn/go-isatty v0.0.12
	github.com/mattn/go-runewidth v0.0.8
	github.com/mattn/go-shellwords v1.0.10
	github.com/milosgajdos83/tenus v0.0.0-20200407174313-f075bef9ab28 // indirect
	github.com/mmcloughlin/avo v0.0.0-20200523190732-4439b6b2c061 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.6.0
	github.com/prometheus/common v0.10.0
	github.com/prometheus/node_exporter v1.0.1
	github.com/pupapaik/sysinfo v0.0.0-20200106202926-c17dea004cd4
	github.com/rs/cors v1.7.0
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980
	google.golang.org/genproto v0.0.0-20200212121238-0849286c0f0e // indirect
	google.golang.org/grpc v1.30.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/xtaci/smux.v1 v1.4.2 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200121175148-a6ecf24a6d71
	modernc.org/httpfs v1.0.0
	nhooyr.io/websocket v1.8.6
)
