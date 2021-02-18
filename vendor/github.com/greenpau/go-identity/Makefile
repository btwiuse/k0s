.PHONY: test ctest covdir coverage docs linter qtest clean dep release logo license
APP_VERSION:=$(shell cat VERSION | head -1)
GIT_COMMIT:=$(shell git describe --dirty --always)
GIT_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD -- | head -1)
BUILD_USER:=$(shell whoami)
BUILD_DATE:=$(shell date +"%Y-%m-%d")
BINARY:="identity"
VERBOSE:=-v
ifdef TEST
	TEST:="-run ${TEST}"
endif

all: test coverage
	@echo "Version: $(APP_VERSION), Branch: $(GIT_BRANCH), Revision: $(GIT_COMMIT)"
	@echo "Build on $(BUILD_DATE) by $(BUILD_USER)"
	@echo "Done!"

linter:
	@echo "Running lint checks"
	@golint *.go
	@echo "PASS: golint"

test: covdir linter
	@go test $(VERBOSE) -coverprofile=.coverage/coverage.out ./*.go

ctest: covdir linter
	@richgo version || go get -u github.com/kyoh86/richgo
	@time richgo test $(VERBOSE) $(TEST) -coverprofile=.coverage/coverage.out ./*.go

covdir:
	@echo "Creating .coverage/ directory"
	@mkdir -p .coverage

coverage:
	@#go tool cover -help
	@go tool cover -html=.coverage/coverage.out -o .coverage/coverage.html
	@go test -covermode=count -coverprofile=.coverage/coverage.out ./*.go
	@go tool cover -func=.coverage/coverage.out | grep -v "100.0"

docs:
	@mkdir -p .doc
	@go doc -all > .doc/index.txt
	@cat .doc/index.txt

clean:
	@rm -rf .doc
	@rm -rf .coverage
	@rm -rf bin/

qtest:
	@echo "Perform quick tests ..."
	@#go test -v -run TestVersioned *.go
	@#go test -v -run TestNewID *.go
	@time richgo test -v -run TestNewPublicKey *.go
	@#time richgo test -v -run TestNewMfaToken *.go
	@#time richgo test -v -run TestNewUser *.go

dep:
	@echo "Making dependencies check ..."
	@go get -u golang.org/x/lint/golint
	@go get -u github.com/greenpau/versioned/cmd/versioned
	@go get -u github.com/google/addlicense
	@go get -u github.com/kyoh86/richgo

license:
	@addlicense -c "Paul Greenberg greenpau@outlook.com" -y 2020 ./*.go

release:
	@echo "Making release"
	@go mod tidy
	@go mod verify
	@if [ $(GIT_BRANCH) != "master" ]; then echo "cannot release to non-master branch $(GIT_BRANCH)" && false; fi
	@git diff-index --quiet HEAD -- || ( echo "git directory is dirty, commit changes first" && false )
	@versioned -patch
	@echo "Patched version"
	@git add VERSION
	@git commit -m 'updated VERSION file'
	@versioned -sync database.go
	@git add database.go
	@git commit -m "released v`cat VERSION | head -1`"
	@git tag -a v`cat VERSION | head -1` -m "v`cat VERSION | head -1`"
	@git push
	@git push --tags
	@#echo "git push --delete origin v$(APP_VERSION)"
	@#echo "git tag --delete v$(APP_VERSION)"

logo:
	@mkdir -p assets/docs/images/
	@gm convert -background black -font Bookman-Demi \
		-size 640x320 "xc:black" \
		-draw "fill white gravity center text 0,0 'Go\nidentity'" \
		-pointsize 96 \
		assets/docs/images/logo.png
