BIN     = k0s
GOOS    = $(shell go env GOOS)
GOARCH  = $(shell go env GOARCH)
LDFLAGS = $(shell ./pkg/version/ldflags)
TAGS    = "nhooyr"

all:
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}"

release:
	@ rm -r releases/latest
	@ go run bin.go -d releases/latest -strip -upx -ldflags="${LDFLAGS}" \
		{linux,android}/{armv6,armv7,arm64,amd64,386} darwin/amd64 windows/{386,amd64}
	@ pushd releases/latest; tree -L 1 -H '.' --noreport --charset utf-8 > index.html; popd
	@ sh -c 'git rev-parse HEAD; git tag -l --points-at HEAD' | \
		xargs -L1 -I@ sh -c 'mkdir -p releases/@; cp -rv releases/latest/* releases/@'
	@ pushd releases; tree -L 1 -H '.' --noreport --charset utf-8 > index.html; popd
	@ git -C releases add .
	@ git -C releases commit -m $(shell git rev-parse HEAD)
	@ git -C releases push

link:
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) bin/$(BIN)
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) bin/agent
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) bin/hub
	ln -f bin/$(BIN)-$(GOOS)-$(GOARCH) bin/client

install:
	install -Dvm755 bin/$(BIN) /usr/bin/$(BIN)
	install -Dvm644 .systemd/agent@.service /usr/lib/systemd/system/agent@.service
	make systemd

systemd:
	@ echo 'Run this manually to initialize/restart the agent systemd service'
	@ echo 'sudo systemctl daemon-reload'
	@ echo 'sudo systemctl enable agent@$$USER'
	@ echo 'sudo systemctl start agent@$$USER'
	@ echo 'sudo systemctl restart agent@$$USER'

clean:
	rm -r bin

up:
	.docker-compose/up

buildkite:
	cd .buildkite && ./gen | tee /dev/stderr > pipeline.yml
