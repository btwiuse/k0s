BIN    = conntroll
GOOS   = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)

all:
	go run bin.go
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) $(BIN)
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) agent
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) hub
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) client

release:
	go run bin.go -strip -upx
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) $(BIN)
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) agent
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) hub
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) client

clean:
	rm -rf bin conntroll agent hub client

