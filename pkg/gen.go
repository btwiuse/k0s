//go:generate protoc --go_out=plugins=grpc:. --proto_path=. msg/types.proto
//go:generate protoc --go_out=plugins=grpc:. --proto_path=. api/api.proto
//go:generate sed -i -e "s,\"msg\",\"github.com/btwiuse/conntroll/pkg/msg\",g" api/api.pb.go
//go:generate goimports -w .

package pkg
