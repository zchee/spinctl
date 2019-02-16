# ----------------------------------------------------------------------------
# global

SHELL = /usr/bin/env bash

GO_PATH ?= $(shell go env GOPATH)
GO_OS ?= $(shell go env GOOS)
GO_ARCH ?= $(shell go env GOARCH)

PKG = $(subst $(GO_PATH)/src/,,$(CURDIR))
GO_PKGS := $(shell go list ./... | grep -v -e '.pb.go' -e 'api/gate')
GO_APP_PKGS := $(shell go list -f '{{if and (or .GoFiles .CgoFiles) (ne .Name "main")}}{{.ImportPath}}{{end}}' ${PKG}/...)
GO_TEST_PKGS := $(shell go list -f='{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}' ./...)
GO_VENDOR_PKGS := $(shell go list -f '{{if and (or .GoFiles .CgoFiles) (ne .Name "main")}}./vendor/{{.ImportPath}}{{end}}' ./vendor/...)

GO_TEST ?= go test
GO_TEST_FUNC ?= .
GO_TEST_FLAGS ?=
GO_BENCH_FUNC ?= .
GO_BENCH_FLAGS ?= -benchmem

VERSION=$(shell cat VERSION.txt)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_UNTRACKED_CHANGES=$(shell git status --porcelain --untracked-files=no)
ifneq ($(GIT_UNTRACKED_CHANGES),)
	GIT_COMMIT := $(GIT_COMMIT)-dirty
endif
CTIMEVAR=-X $(PKG)/pkg/version.gitCommit=$(GIT_COMMIT) -X $(PKG)/pkg/version.version=$(VERSION)

CGO_ENABLED ?= 0
GO_LDFLAGS=-s -w $(CTIMEVAR)
GO_LDFLAGS_STATIC=-s -w '-extldflags=-static' $(CTIMEVAR)
ifneq ($(GO_OS),darwin)
	GO_LDFLAGS_STATIC+=-d
endif

ifneq ($(wildcard go.mod),)  # exist go.mod
ifeq ($(GO111MODULE),off)
	GO_FLAGS=-mod=vendor
endif
endif

GO_BUILDTAGS := osusergo netgo
GO_FLAGS ?= -tags='$(GO_BUILDTAGS)' -ldflags="${GO_LDFLAGS}"
GO_INSTALLSUFFIX_STATIC=netgo
GO_BUILDTAGS_STATIC=static static_build

CONTAINER_REGISTRY := gcr.io/container-image

# ----------------------------------------------------------------------------
# defines

GOPHER=""
define target
@printf "$(GOPHER)  \\x1b[1;32m$(patsubst ,$@,$(1))\\x1b[0m\\n"
endef

# ----------------------------------------------------------------------------
# targets

## build and install

.PHONY: $(APP)
$(APP): VERSION.txt
	$(call target)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GO_OS) GOARCH=$(GO_ARCH) go build -v $(strip $(GO_FLAGS)) -o $(APP) $(CMD)

.PHONY: build
build: GO_FLAGS+=-ldflags="${GO_LDFLAGS}"
build: $(APP)  ## Builds a dynamic executable or package.

.PHONY: static
static: GO_FLAGS+=-ldflags="${GO_LDFLAGS_STATIC}" -installsuffix ${GO_INSTALLSUFFIX_STATIC}
static: GO_BUILDTAGS+=${GO_BUILDTAGS_STATIC}
static: $(APP)  ## Builds a static executable or package.

.PHONY: install
install: GO_FLAGS+=-a -ldflags="${GO_LDFLAGS_STATIC}" -installsuffix ${GO_INSTALLSUFFIX_STATIC}
install: GO_BUILDTAGS+=${GO_BUILDTAGS_STATIC}
install:  ## Installs the executable or package.
	$(call target)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GO_OS) GOARCH=$(GO_ARCH) go install -v $(strip $(GO_FLAGS)) $(CMD)

.PHONY: pkg/install
pkg/install: GO_FLAGS+=-ldflags="${GO_LDFLAGS_STATIC}" -installsuffix ${GO_INSTALLSUFFIX_STATIC}
pkg/install: GO_BUILDTAGS+=${GO_BUILDTAGS_STATIC}
pkg/install:
	$(call target)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GO_OS) GOARCH=$(GO_ARCH) go install -v $(strip $(GO_FLAGS)) ${GO_APP_PKGS}


## test, bench and coverage

.PHONY: test
test:  ## Run the package test with checks race condition.
	$(call target)
	$(GO_TEST) -v -race $(strip $(GO_FLAGS)) -run=$(GO_TEST_FUNC) $(GO_TEST_PKGS)

.PHONY: bench
bench:  ## Take a package benchmark.
	$(call target)
	$(GO_TEST) -v $(strip $(GO_FLAGS)) -run='^$$' -bench=$(GO_BENCH_FUNC) -benchmem $(GO_TEST_PKGS)

.PHONY: bench/race
bench/race:  ## Take a package benchmark with checks race condition.
	$(call target)
	$(GO_TEST) -v -race $(strip $(GO_FLAGS)) -run='^$$' -bench=$(GO_BENCH_FUNC) -benchmem $(GO_TEST_PKGS)

.PHONY: bench/trace
bench/trace:  ## Take a package benchmark with take a trace profiling.
	$(GO_TEST) -v -c -o bench-trace.test $(PKG)/stackdriver
	GODEBUG=allocfreetrace=1 ./bench-trace.test -test.run=none -test.bench=$(GO_BENCH_FUNC) -test.benchmem -test.benchtime=10ms 2> trace.log

.PHONY: coverage
coverage: GO_FLAGS+=-ldflags="${GO_LDFLAGS_STATIC}" -installsuffix ${GO_INSTALLSUFFIX_STATIC}
coverage: GO_BUILDTAGS+=${GO_BUILDTAGS_STATIC}
coverage: clean  ## Take test coverage.
	$(call target)
	@$(GO_TEST) -v -race $(strip $(GO_FLAGS)) -covermode=atomic -coverpkg=$(PKG)/pkg/... -coverprofile=coverage.out $(GO_PKGS)

$(GO_PATH)/bin/go-junit-report:
	@GO111MODULE=off go get -u github.com/jstemmer/go-junit-report

.PHONY: cmd/go-junit-report
cmd/go-junit-report: $(GO_PATH)/bin/go-junit-report  # go get 'go-junit-report' binary

.PHONY: coverage/ci
coverage/ci: GO_FLAGS+=-a -ldflags="${GO_LDFLAGS_STATIC}" -installsuffix ${GO_INSTALLSUFFIX_STATIC}
coverage/ci: GO_BUILDTAGS+=${GO_BUILDTAGS_STATIC}
coverage/ci: cmd/go-junit-report  ## Take test coverage.
	$(call target)
	@mkdir -p /tmp/ci/artifacts /tmp/ci/test-results
	$(GO_TEST) -v -race $(strip $(GO_FLAGS)) -covermode=atomic -coverpkg=$(PKG)/pkg/... -coverprofile=/tmp/ci/artifacts/coverage.out $(GO_PKGS) 2>&1 | tee /dev/stderr | go-junit-report -set-exit-code > /tmp/ci/test-results/junit.xml
	@go tool cover -html=/tmp/ci/artifacts/coverage.out -o /tmp/ci/artifacts/coverage.html


## lint

.PHONY: lint
lint: lint/golangci-lint  ## Run all linters.

$(GO_PATH)/bin/golangci-lint:
	GO111MODULE=off go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: cmd/golangci-lint
cmd/golangci-lint: $(GO_PATH)/bin/golangci-lint  # go get 'golangci-lint' binary

.PHONY: golangci-lint
lint/golangci-lint: cmd/golangci-lint .golangci.yml  ## Run golangci-lint.
	$(call target)
	@GO111MODULE=on golangci-lint run ./...


## mod

.PHONY: mod/init
mod/init:
	$(call target)
	@GO111MODULE=on go mod init > /dev/null 2>&1 || true

.PHONY: mod/goget
mod/goget:  ## Update module and go.mod.
	$(call target)
	@GO111MODULE=on go get -u -m -v -x ./...

.PHONY: mod/tidy
mod/tidy:
	$(call target)
	@GO111MODULE=on go mod tidy -v

.PHONY: mod/vendor
mod/vendor: mod/tidy
	$(call target)
	@GO111MODULE=on go mod vendor -v

.PHONY: mod/graph
mod/graph:
	$(call target)
	@GO111MODULE=on go mod graph

.PHONY: mod/clean
mod/clean:
	$(call target)
	@$(RM) -r $(shell find vendor -maxdepth 1 -path "vendor/*" -type d)

.PHONY: mod/install
mod/install: GO_FLAGS+=-ldflags="${GO_LDFLAGS_STATIC}" -installsuffix netgo
mod/install: GO_BUILDTAGS+=netgo static static_build
mod/install: mod/tidy mod/vendor
	$(call target)
	@GO111MODULE=off go install -v $(strip $(GO_FLAGS)) $(GO_VENDOR_PKGS) || GO111MODULE=on go install -mod=vendor -v $(strip $(GO_FLAGS)) $(GO_VENDOR_PKGS)

.PHONY: mod/update
mod/update: mod/goget mod/tidy mod/vendor mod/install  ## Updates all vendor packages.

.PHONY: mod
mod: mod/tidy mod/vendor mod/install  ## Updates the vendoring directory via go mod.
	@sed -i ':a;N;$$!ba;s|go 1\.12\n\n||g' go.mod


## miscellaneous

.PHONY: container/build
container/build:  ## Create the container image from the Dockerfile.
	docker image build --rm --force-rm --pull --progress=plain -t $(CONTAINER_REGISTRY)/$(APP):$(VERSION:v%=%) .

.PHONY: container/push
container/push:  ## Push the container image to $CONTAINER_REGISTRY
	docker image push $(CONTAINER_REGISTRY)/$(APP):$(VERSION:v%=%)


.PHONY: boilerplate/go/%
boilerplate/go/%: BOILERPLATE_PKG_DIR=$(shell printf $@ | cut -d'/' -f3- | rev | cut -d'/' -f2- | rev)
boilerplate/go/%: BOILERPLATE_PKG_NAME=$(if $(findstring $@,cmd),main,$(shell printf $@ | rev | cut -d/ -f2 | rev))
boilerplate/go/%: hack/boilerplate/boilerplate.go.txt  ## Create go file from boilerplate.go.txt
	@if [ ! -d ${BOILERPLATE_PKG_DIR} ]; then mkdir -p ${BOILERPLATE_PKG_DIR}; fi
	@cat hack/boilerplate/boilerplate.go.txt <(printf "package ${BOILERPLATE_PKG_NAME}\\n") > $*
	@sed -i "s|YEAR|$(shell date '+%Y')|g" $*


.PHONY: AUTHORS
AUTHORS:  ## Creates AUTHORS file.
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@printf "$(shell git log --format="\n%aN <%aE>" | LC_ALL=C.UTF-8 sort -uf)" >> $@


.PHONY: clean
clean:  ## Cleanup any build binaries or packages.
	$(call target)
	@$(RM) $(APP) *.out *.test *.prof trace.log


.PHONY: help
help:  ## Show make target help.
	@perl -nle 'BEGIN {printf "Usage:\n  make \033[33m<target>\033[0m\n\nTargets:\n"} printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 if /^([a-zA-Z\/_-].+)+:.*?\s+## (.*)/' ${MAKEFILE_LIST}
