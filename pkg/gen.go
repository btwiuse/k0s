//go:generate protoc --go_out=plugins=grpc:. --proto_path=. msg/types.proto
//go:generate cp ./github.com/btwiuse/wetty/pkg/msg/types.pb.go msg/types.pb.go
//go:generate rm -r github.com
//go:generate protoc --go_out=plugins=grpc:. --proto_path=. api/api.proto
//go:generate goimports -w .

package pkg
