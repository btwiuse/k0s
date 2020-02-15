NAME     := k0s
PACKAGE  := github.com/btwiuse/$(NAME)
GIT      := $(shell git rev-parse --short HEAD)
DATE     := $(shell date +%FT%T%Z)
VERSION  := v0.0.2
IMG_NAME := btwiuse/k0s
IMAGE    := ${IMG_NAME}:${VERSION}
LDFLAGS  := $(shell ./pkg/version/ldflags)
TAGS     := ""
SHELL    := bash

default: help

fonts:
	@ mkdir -p fonts/ pkg/fonts/; cp /usr/share/figlet/fonts/standard.flf fonts/
	@ assets -d fonts/ -package fonts -o ./pkg/fonts/standard.go -map Fonts
	@ rm -r fonts/

gazelle:             ## auto generate BUILD.bazel files from go.mod
	@ go mod tidy
	@ go mod vendor
	@ bazel run //:gazelle -- update-repos --from_file=go.mod
	@ bazel run //:gazelle
	# @ gazelle -exclude=vendor
	# @ git checkout vendor/nhooyr.io/websocket
	@ git status vendor/

bazel-build-android:            ## Build android binaries using bazel
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:android_amd64  //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:android_386    //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:android_arm    //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:android_arm64  //:k0s

bazel-build-windows:            ## Build windows binaries using bazel
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:windows_amd64  //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:windows_386    //:k0s

bazel-build-darwin:            ## Build darwin binaries using bazel
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64  //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:darwin_386    //:k0s

bazel-build-bsd:            ## Build bsd binaries using bazel
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:freebsd_amd64  //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:freebsd_386    //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:openbsd_amd64  //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:openbsd_386    //:k0s

bazel-build-bsd-arm:            ## Build bsd arm binaries using bazel
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:freebsd_arm    //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:openbsd_arm    //:k0s

bazel-build-linux-arm:          ## Build linux arm binaries using bazel
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_arm64    //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_arm      //:k0s

bazel-build-linux:          ## Build linux binaries using bazel
	@ bazel build //:k0s # //cmd/{hub,client,agent}
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64    //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_arm64    //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_386      //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_mips     //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_mips64   //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_mipsle   //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_mips64le //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_ppc64    //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_ppc64le  //:k0s
	@ bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_s390x    //:k0s

bazel-build:          ## Build binary for current platform using bazel
	@ bazel build //:k0s # //cmd/{hub,client,agent}

build:          ## Build binary for current platform
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}"

dry:      ## Build binary for every supported platform
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}" -dry \
	  linux/{{mips{,64},ppc64}{,le},s390x} \
		{linux,android}/{armv6,armv7,arm64,amd64,386} {darwin,windows}/{386,amd64} \
		freebsd/{armv7,armv6} {freebsd,openbsd}/{386,amd64}

build-all:      ## Build binary for every supported platform
	@ make build-android
	@ make build-linux
	@ make build-linux-arm
	@ make build-bsd
	@ make build-bsd-arm
	@ make build-windows
	@ make build-darwin

build-android:  ## Build android binaries
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
		android/{armv6,armv7,arm64,amd64,386}

build-linux-arm: ## Build linux arm binaries
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
		linux/{armv6,armv7,arm64}

build-bsd-arm:  	## Build bsd arm binaries
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
	  freebsd/{armv7,armv6} # ,arm64

build-bsd:  	## Build bsd binaries
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
	  {freebsd,openbsd}/{amd64,386}

build-linux:  	## Build linux binaries
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
	  linux/{amd64,386} linux/{{mips{,64},ppc64}{,le},s390x}

build-windows:  ## Build windows binaries
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
		windows/{386,amd64}

build-darwin:   ## Build darwin binaries
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
		darwin/{386,amd64}

scratch-build:  ## Build without using existing build cache
	@ go run bin.go -d releases/latest -ldflags="${LDFLAGS}" -- -a

scratch-build-all:      ## Build binary for every supported platform ignoring build cache
	@ go run bin.go -tags "$(TAGS)" -ldflags="${LDFLAGS}" -- -a\
		{linux,android}/{armv6,armv7,arm64,amd64,386} {darwin,windows}/{386,amd64} \
	  linux/{{mips{,64},ppc64}{,le},s390x}

release:        ## Build and upload binaries for all supported platforms
	@ mkdir -p bin/; git -C bin/ init
	@ make build-all
	@ pushd bin; tree -L 1 -H '.' --noreport --charset utf-8 > index.html; popd
	@ .ci/release-latest.sh

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
