module k0s.io/pkg/tunnel

go 1.19

replace (
	k0s.io => ../../
	k0s.io/pkg/wrap => ../wrap
)

require (
	github.com/gorilla/handlers v1.5.1
	github.com/rs/cors v1.8.2
	k0s.io/pkg/wrap v0.0.0-00010101000000-000000000000
	nhooyr.io/websocket v1.8.7
)

require (
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/klauspost/compress v1.10.3 // indirect
)
