.PHONY: build trust

NAME     := k0s
PACKAGE  := github.com/btwiuse/$(NAME)
GIT      := $(shell git rev-parse --short HEAD)
DATE     := $(shell date +%FT%T%Z)
VERSION  := $(shell cat VERSION)
IMG_NAME := btwiuse/k0s
IMAGE    := ${IMG_NAME}
LDFLAGS  := $(shell ./tools/ldflags)
TAGS     := $(TAGS)
SHELL    := bash
BAZEL    := $(shell ./tools/which_bazel)
BINGO    := go tool github.com/btwiuse/bingo/cmd/bingo

default: help

goreleaser:
	@ GORELEASER_CURRENT_TAG=v0.0.0 goreleaser release --snapshot --clean --parallelism 1 --timeout 1h

docker-login:
	@ docker login -u $(DOCKERHUB_USERNAME) -p $(DOCKERHUB_TOKEN)

devcontainer: docker-login
	@ docker build -t btwiuse/k0s:devcontainer -f .devcontainer/Dockerfile .devcontainer
	@ docker push btwiuse/k0s:devcontainer

fonts:
	@ mkdir -p fonts/ pkg/fonts/; cp /usr/share/figlet/fonts/standard.flf fonts/
	@ assets -d fonts/ -package fonts -o ./pkg/fonts/standard.go -map Fonts
	@ rm -r fonts/

raze:             ## auto generate BUILD.bazel files from Cargo.toml
	@ which cargo-raze || (echo please cargo install cargo-raze | grep --color . 1>&2 && false)
	@#cargo vendor --versioned-dirs cargo/vendor &>/dev/null
	@ cargo raze
	# @ $(BAZEL) run @cargo_raze//:raze -- --manifest-path=$(realpath /Cargo.toml)
	@ git status cargo

replace:         ## inject replace directives to all workspace modules based on go.work
	@ ./tools/replace_all

require:         ## turn replace directives for workspace modules into require
	@ ./tools/require_all

work-sync:
	go work sync

gazelle: work-sync             ## auto generate BUILD.bazel files from go.mod
	@ go mod tidy
	@ $(BAZEL) mod tidy
	@#go mod vendor
	@ find pkg -name 'go.mod' | sed s,go.mod,,g | xargs -I% bash -vc 'pushd % && go mod tidy'
	@#sed -i -e '/k0s.io.* v/d' pkg/*/go.mod go.mod
	@#ls -1 pkg/*/go.mod | xargs -L1 $(BAZEL) run //:gazelle -- update-repos --from_file
	@# @ $(BAZEL) run //:gazelle -- update-repos --from_file=third_party/go.mod -lang go -proto disable_global
	@# https://github.com/btwiuse/baize/blob/btwiuse/hack/update-deps.sh
	@# $(BAZEL) run //:gazelle -- update-repos --from_file=third_party/go.mod --build_file_generation=on --build_file_proto_mode=disable --prune --to_macro=go_repos.bzl%go_repositories
	@ $(BAZEL) run //:gazelle -- -build_file_name=BUILD.bazel
	@#git status vendor/

cancel-actions-in-queue:     ## cancel gh actions in queue
	@ gh run list -L 60 | grep queued | cut -f 7 | xargs -L1 -P8 gh run cancel

cancel-actions:     ## cancel gh actions
	@ gh run list -L 60 | cut -f 7 | xargs -L1 -P8 gh run cancel

delete-releases:    ## delete all non-version releases
	@ gh release list -L 100 | grep ^cmd | cut -f 1 | xargs -L1 -P8 gh release delete --yes
	@ gh release list -L 100 | grep ^pkg | cut -f 1 | xargs -L1 -P8 gh release delete --yes

go-get:               ## trigger update for https://pkg.go.dev
	@ go work edit --json | grep ModPath | grep -o '"k0s.io/.*"' | xargs -I% bash -c 'echo go get %@$(shell git rev-parse HEAD)' | bash -v

go-install:           ## install latest commit from https://pkg.go.dev
	@ go install -v k0s.io/cmd/k0s@$(shell git rev-parse HEAD)

go-install-debuginfo:     ## install latest commit from https://pkg.go.dev with debuginfo
	@ go install -tags runtime_debug_buildinfo -v k0s.io/cmd/k0s@$(shell git rev-parse HEAD)
	@ k0s hub -version

bazel-info:                ## Bazel info (buildkite precheck)
	$(BAZEL) --nosystem_rc --nohome_rc info

bazel-build-android:            ## Build android binaries using bazel
	$(BAZEL) run //:install_k0s --config=go_android_amd64 -- -g $(PWD)/bin/android/amd64
	$(BAZEL) run //:install_k0s --config=go_android_386   -- -g $(PWD)/bin/android/386
	$(BAZEL) run //:install_k0s --config=go_android_armv7 -- -g $(PWD)/bin/android/armv7
	$(BAZEL) run //:install_k0s --config=go_android_arm64 -- -g $(PWD)/bin/android/arm64
	# @ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:android_amd64  //:k0s
	# @ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:android_386    //:k0s
	# @ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:android_arm    //:k0s
	# @ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:android_arm64  //:k0s
	# $(BAZEL) build //:k0s --config=go_android_amd64
	# $(BAZEL) build //:k0s --config=go_android_386
	# $(BAZEL) build //:k0s --config=go_android_armv7
	# $(BAZEL) build //:k0s --config=go_android_arm64

bazel-build-windows:            ## Build windows binaries using bazel
	$(BAZEL) run --platforms=@io_bazel_rules_go//go/toolchain:windows_amd64 //:install_k0s_static -- -g $(PWD)/bin/windows/amd64
	$(BAZEL) run --platforms=@io_bazel_rules_go//go/toolchain:windows_386 //:install_k0s_static -- -g $(PWD)/bin/windows/386
	$(BAZEL) run --platforms=@io_bazel_rules_go//go/toolchain:windows_arm //:install_k0s_static -- -g $(PWD)/bin/windows/armv7
	$(BAZEL) run --platforms=@io_bazel_rules_go//go/toolchain:windows_arm64 //:install_k0s_static -- -g $(PWD)/bin/windows/arm64
	# $(BAZEL) build //:k0s --config=go_win32
	# $(BAZEL) build //:k0s --config=go_win64

bazel-build-darwin:            ## Build darwin binaries using bazel
	$(BAZEL) run --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64  //:install_k0s -- -g $(PWD)/bin/darwin/amd64
	$(BAZEL) run --platforms=@io_bazel_rules_go//go/toolchain:darwin_arm64  //:install_k0s -- -g $(PWD)/bin/darwin/arm64
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64  //:k0s
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:darwin_arm64  //:k0s

bazel-build-bsd:            ## Build bsd binaries using bazel
	@ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:freebsd_amd64  //:k0s
	@ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:freebsd_386    //:k0s
	@ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:openbsd_amd64  //:k0s
	@ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:openbsd_386    //:k0s

bazel-build-bsd-arm:            ## Build bsd arm binaries using bazel
	@ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:freebsd_arm    //:k0s
	@ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:openbsd_arm    //:k0s

bazel-build-linux-arm:          ## Build linux arm binaries using bazel
	@ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_arm64    //:k0s_static
	@ $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_arm      //:k0s_static

bazel-build-linux:          ## Build linux binaries using bazel
	$(BAZEL) run   --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64    //:install_k0s_static -- -g $(PWD)/bin/linux/amd64
	$(BAZEL) run   --platforms=@io_bazel_rules_go//go/toolchain:linux_386      //:install_k0s_static -- -g $(PWD)/bin/linux/386
	$(BAZEL) run   --platforms=@io_bazel_rules_go//go/toolchain:linux_arm64    //:install_k0s_static -- -g $(PWD)/bin/linux/arm64
	$(BAZEL) run   --platforms=@io_bazel_rules_go//go/toolchain:linux_arm      //:install_k0s_static -- -g $(PWD)/bin/linux/arm
	# $(BAZEL) run   --platforms=@io_bazel_rules_go//go/toolchain:linux_mips     //:install_k0s_static -- -g $(PWD)/bin/linux/mips
	# $(BAZEL) run   --platforms=@io_bazel_rules_go//go/toolchain:linux_mips64   //:install_k0s_static -- -g $(PWD)/bin/linux/mips64
	# $(BAZEL) run   --platforms=@io_bazel_rules_go//go/toolchain:linux_mipsle   //:install_k0s_static -- -g $(PWD)/bin/linux/mipsle
	# $(BAZEL) run   --platforms=@io_bazel_rules_go//go/toolchain:linux_mips64le //:install_k0s_static -- -g $(PWD)/bin/linux/mips64le
	# $(BAZEL) run   --platforms=@io_bazel_rules_go//go/toolchain:linux_ppc64    //:install_k0s_static -- -g $(PWD)/bin/linux/ppc64
	# $(BAZEL) run   --platforms=@io_bazel_rules_go//go/toolchain:linux_ppc64le  //:install_k0s_static -- -g $(PWD)/bin/linux/ppc64le
	# $(BAZEL) run   --platforms=@io_bazel_rules_go//go/toolchain:linux_s390x    //:install_k0s_static -- -g $(PWD)/bin/linux/s390x
	# $(BAZEL) build //:k0s_static # //cmd/{hub,client,agent}
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64    //:k0s_static
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_arm64    //:k0s_static
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_386      //:k0s_static
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_mips     //:k0s_static
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_mips64   //:k0s_static
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_mipsle   //:k0s_static
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_mips64le //:k0s_static
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_ppc64    //:k0s_static
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_ppc64le  //:k0s_static
	# $(BAZEL) build --platforms=@io_bazel_rules_go//go/toolchain:linux_s390x    //:k0s_static

bazel-build:          ## Build binary for current platform using bazel
	$(BAZEL) build //:k0s
	# $(BAZEL) run //:install_k0s_static -- -g $(PWD)/bin
	# $(BAZEL) build # //:k0s # //cmd/{hub,client,agent}

build:          ## Build binary for current platform
	@ $(BINGO) -tags "$(TAGS)" -ldflags="${LDFLAGS}"

dry:      ## Build binary for every supported platform
	@ $(BINGO) -tags "$(TAGS)" -ldflags="${LDFLAGS}" -dry \
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
	@ $(BINGO) -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
		android/{armv6,armv7,arm64,amd64,386}

build-bsd-arm:  	## Build bsd arm binaries
	@ $(BINGO) -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
	  {netbsd,openbsd,freebsd}/{armv6,armv7,arm64}

build-bsd:  	## Build bsd binaries
	@ $(BINGO) -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
	  {netbsd,openbsd,freebsd}/{amd64,386}

build-linux:  	## Build linux binaries
	@ $(BINGO) -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
	  linux/{amd64,386}

build-linux-arm: ## Build linux arm binaries
	@ $(BINGO) -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
	  linux/{armv6,armv7,arm64}

build-linux-others:  	## Build linux binaries
	@ $(BINGO) -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
	  linux/{{mips{,64},ppc64}{,le},s390x,{riscv,loong}64}

build-windows:  ## Build windows binaries
	@ $(BINGO) -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
		windows/{386,amd64,armv7,arm64}

build-darwin:   ## Build darwin binaries
	@ $(BINGO) -tags "$(TAGS)" -ldflags="${LDFLAGS}" \
		darwin/{amd64,arm64}

scratch-build:  ## Build without using existing build cache
	@ $(BINGO) -d releases/latest -ldflags="${LDFLAGS}" -- -a

scratch-build-all:      ## Build binary for every supported platform ignoring build cache
	@ $(BINGO) -tags "$(TAGS)" -ldflags="${LDFLAGS}" -- -a\
		{linux,android}/{armv6,armv7,arm64,amd64,386} {darwin,windows}/{386,amd64} \
	  linux/{{mips{,64},ppc64}{,le},s390x}

release:        ## Build and upload binaries for all supported platforms
	# mkdir -p bin/; git -C bin/ init
	# make build
	# pushd bin; tree -L 1 -H '.' --noreport --charset utf-8 > index.html; popd
	@ .ci/release-latest.sh

dist:           ## Build and make an dist image TODO: android builder image
	@ make build
	@ make build-linux-arm
	@ make build-windows
	@ make build-darwin
	@ pushd bin; test -f index.html && rm index.html; tree -H '.' --noreport --charset utf-8 | sponge index.html; popd
	@ cp VERSION bin/
	@ ./bin/k0s mnt https://chassis.k0s.io/btwiuse/k0s/v0.0.11/ ./bin/

install:        ## install binary to system paths
	install -Dvm755 bin/$(NAME) /usr/bin/$(NAME)
	install -Dvm644 .systemd/k0s-agent@.service /etc/systemd/system/k0s-agent@.service
	install -Dvm644 .systemd/k0s-hub.service /etc/systemd/system/k0s-hub.service
	make systemd

systemd:        ## Show systemd post install actions
	@ echo
	@ echo '# Reload systemd unit files'
	@ echo 'sudo systemctl daemon-reload'
	@ echo
	@ echo '# Run this manually to initialize/restart the agent service'
	@ echo 'sudo systemctl enable k0s-agent@$$USER'
	@ echo 'sudo systemctl start k0s-agent@$$USER'
	@ echo 'sudo systemctl restart k0s-agent@$$USER'
	@ echo
	@ echo '# Run this manually to initialize/restart the hub service'
	@ echo 'sudo systemctl enable k0s-hub'
	@ echo 'sudo systemctl start k0s-hub'
	@ echo 'sudo systemctl restart k0s-hub'
	@ echo
	@ echo '# To stop those services，do the following:'
	@ echo 'sudo systemctl list-units | grep k0s'
	@ echo 'sudo systemctl stop k0s-agent@$$USER'
	@ echo 'sudo systemctl stop k0s-hub'

clean:          ## Clean build artifacts
	rm -r bin

up:             ## Start prometheus and grafana to scrape and display metrics
	.docker-compose/up

buildkite:      ## Generate buildkite pipeline yml definition
	cd .buildkite && ./gen | tee /dev/stderr > pipeline.yml

test-build:           ## Check all build tags will compile
	@ $(BINGO) -tags raw        # test tag raw
	@ $(BINGO) -tags nhooyr     # test tag nhooyr
	@ $(BINGO) -tags gorilla    # test tag gorilla

test:           ## Run all tests
	@go clean --testcache && go test ./...


cover:          ## Run test coverage suite
	@go test ./... --coverprofile=cov.out
	@go tool cover --html=cov.out

img:  docker-login   ## Build Docker Image and push
	@ docker build --rm -t ${IMAGE} -f Dockerfile .
	@ docker push ${IMAGE}

bazel-img:      ## Build Docker Image archive with bazel
	@ ${BAZEL} build //:k0s_image.tar

bazel-docker-img: bazel-img      ## Build Docker Image with bazel
	@ docker load -i bazel-bin/k0s_image.tar
	@ docker tag bazel:k0s_image ${IMAGE}

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'
