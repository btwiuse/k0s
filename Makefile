NAME     := k0s
PACKAGE  := github.com/btwiuse/$(NAME)
GIT      := $(shell git rev-parse --short HEAD)
DATE     := $(shell date +%FT%T%Z)
VERSION  := v0.0.2
IMG_NAME := btwiuse/k0s
IMAGE    := ${IMG_NAME}:${VERSION}
LDFLAGS  := $(shell ./pkg/version/ldflags)
TAGS     := ""

default: help

build:          ## Build binary for current platform
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}"

release:        ## Build and upload binaries for all supported platforms
	@ mkdir -p releases/latest ; git -C releases init; rm -r releases/latest
	@ go run bin.go -d releases/latest -strip -upx -ldflags="${LDFLAGS}" \
		{linux,android}/{armv6,armv7,arm64,amd64,386} {darwin,windows}/{386,amd64} \
	  linux/{{mips{,64},ppc64}{,le},s390x}
	@ pushd releases/latest; tree -H '.' --noreport --charset utf-8 > index.html; popd
	@ sh -c 'git rev-parse HEAD; git tag -l --points-at HEAD' | \
		xargs -L1 -I@ sh -c 'mkdir -p releases/@; cp -rv releases/latest/* releases/@'
	@ pushd releases; tree -L 1 -H '.' --noreport --charset utf-8 > index.html; popd
	@ git -C releases add .
	@ git -C releases commit -m $(shell git rev-parse HEAD)
	@ git -C releases push

latest:         ## update latest tag
	@ git push --delete origin latest || true
	@ git tag -d latest || true
	@ git tag latest HEAD
	@ git push origin latest

install:        ## install binary to system paths
	install -Dvm755 bin/$(NAME) /usr/bin/$(NAME)
	install -Dvm644 .systemd/agent@.service /usr/lib/systemd/system/agent@.service
	make systemd

systemd:        ## Show systemd post install actions
	@ echo 'Run this manually to initialize/restart the agent systemd service'
	@ echo 'sudo systemctl daemon-reload'
	@ echo 'sudo systemctl enable agent@$$USER'
	@ echo 'sudo systemctl start agent@$$USER'
	@ echo 'sudo systemctl restart agent@$$USER'

clean:          ## Clean build artifacts
	rm -r bin

up:             ## Start prometheus and grafana to scrape and display metrics
	.docker-compose/up

buildkite:      ## Generate buildkite pipeline yml definition
	cd .buildkite && ./gen | tee /dev/stderr > pipeline.yml

test-build:           ## Check all build tags will compile
	@ go run bin.go -tags raw        # test tag raw
	@ go run bin.go -tags nhooyr     # test tag nhooyr
	@ go run bin.go -tags gorilla    # test tag gorilla

test:           ## Run all tests
	@go clean --testcache && go test ./...


cover:          ## Run test coverage suite
	@go test ./... --coverprofile=cov.out
	@go tool cover --html=cov.out

img:            ## Build Docker Image
	@docker build --rm -t ${IMAGE} .

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'
