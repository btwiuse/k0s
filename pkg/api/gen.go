//go:generate protoc --go_out=plugins=grpc:. api.proto
//go:generate goimports -w .

package api
