module k0s.io/k0s

go 1.14

// replace github.com/btwiuse/wetty => /home/aaron/wetty
// replace github.com/yrpc/yrpc => /home/aaron/yrpc/yrpc
replace nhooyr.io/websocket => github.com/btwiuse/websocket v1.7.5-0.20200215101310-b4b58f88e17f

require (
	github.com/MaxRis/w32 v0.0.0-20180517000239-4f5cfb03fabf // indirect
	github.com/VojtechVitek/yaml-cli v0.0.5
	github.com/abbot/go-http-auth v0.4.0
	github.com/alexpantyukhin/go-pattern-match v0.0.0-20200207085858-033d90da7458
	github.com/btwiuse/gods v0.0.0-20190414062120-7e7cf0aebbb0
	github.com/btwiuse/pretty v0.0.0-20190401073227-519ff4ea1882
	github.com/btwiuse/wetty v0.0.33
	github.com/creack/pty v1.1.9
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/docker/docker v1.4.2-0.20200214221943-d8772509d1a2
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/gdamore/tcell v1.3.0
	github.com/golang/protobuf v1.4.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/websocket v1.4.1
	github.com/kr/pty v1.1.8 // indirect
	github.com/liamg/aminal v0.9.0
	github.com/lucasb-eyer/go-colorful v1.0.3 // indirect
	github.com/lukesampson/figlet v0.0.0-20190211215653-8a3ef4a6ac42
	github.com/mattn/go-isatty v0.0.12
	github.com/mattn/go-runewidth v0.0.8
	github.com/mattn/go-shellwords v1.0.10
	github.com/mdp/qrterminal v1.0.1
	github.com/miekg/dns v1.1.27 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.6.0
	github.com/prometheus/common v0.10.0
	github.com/prometheus/node_exporter v1.0.1
	github.com/pupapaik/sysinfo v0.0.0-20200106202926-c17dea004cd4
	github.com/riobard/go-shadowsocks2 v0.2.1
	github.com/rs/cors v1.7.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/txthinking/gotun2socks v0.0.0-20180829122610-35016fdae05e
	github.com/txthinking/socks5 v0.0.0-20190404052647-254e122c4eaf
	github.com/txthinking/x v0.0.0-20190708114625-99b19c1440b6
	github.com/urfave/negroni v1.0.0
	github.com/yrpc/yrpc v0.0.0-20191231114812-451503bf48c2
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980
	google.golang.org/genproto v0.0.0-20200212121238-0849286c0f0e // indirect
	google.golang.org/grpc v1.30.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/russross/blackfriday.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.3.0
	gopkg.in/yaml.v3 v3.0.0-20200121175148-a6ecf24a6d71 // indirect
	modernc.org/httpfs v1.0.0
	nhooyr.io/websocket v1.7.5-0.20200215101310-b4b58f88e17f
)
