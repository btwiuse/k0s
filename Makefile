BIN    = conntroll
GOOS   = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)

all:
	go run bin.go -tags "$(TAGS)"

release:
	go run bin.go -d releases/$(shell git rev-parse HEAD) -strip -upx linux/{arm,arm64,amd64,386} darwin/amd64

link:
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) bin/$(BIN)
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) bin/agent
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) bin/hub
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) bin/client

install:
	install -Dvm755 bin/$(BIN) /usr/bin/$(BIN)
	ln -f /usr/bin/$(BIN) /usr/bin/agent
	@# ln -f /usr/bin/$(BIN) /usr/bin/hub
	@# ln -f /usr/bin/$(BIN) /usr/bin/client
	install -Dvm644 agent@.service /usr/lib/systemd/system/agent@.service
	@ echo "Now manually run:"
	@ echo 'sudo systemctl daemon-reload'
	@ echo 'sudo systemctl restart agent@$$USER'
	@ echo 'or, if you are installing for the first time'
	@ echo 'sudo systemctl enable agent@$$USER'
	@ echo 'sudo systemctl start agent@$$USER'

clean:
	rm -r bin
