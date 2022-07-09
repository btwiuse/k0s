module k0s.io/pkg/api

go 1.19

replace (
	k0s.io => ../../
	k0s.io/pkg/agent => ../agent/
)

require (
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.47.0
	k0s.io/pkg/agent v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/sys v0.0.0-20220708085239-5a0f0661e09d // indirect
	golang.org/x/text v0.3.8-0.20211004125949-5bd84dd9b33b // indirect
	google.golang.org/genproto v0.0.0-20220706185917-7780775163c4 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	k0s.io v0.0.0-00010101000000-000000000000 // indirect
)
