BIN    = conntroll
GOOS   = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)

all:
	go run bin.go
	ln -sf bin/$(BIN)-$(GOOS)-$(GOARCH) $(BIN)

release:
	go run bin.go -strip -upx

clean:
	rm -rf bin conntroll

