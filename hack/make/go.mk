# ----------------------------------------------------------------------------
# global

SHELL = /usr/bin/env bash
 
GO_PATH := $(shell go env GOPATH)
PKG = $(subst $(GO_PATH)/src/,,$(CURDIR))
GO_PKGS := $(shell go list ./... | grep -v -e '.pb.go' -e 'api/gate' -e 'api/mock')
GO_PKGS_ABS := $(shell go list -f '$(GO_PATH)/src/{{.ImportPath}}' ./... | grep -v -e '.pb.go' -e 'api/gate' -e 'api/mock')
GO_TEST_PKGS := $(shell go list -f='{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}' ./...)
GO_VENDOR_PKGS := $(shell go list -f '{{if and (or .GoFiles .CgoFiles) (ne .Name "main")}}{{.ImportPath}}{{end}}' ./vendor/...)

VERSION := $(shell cat VERSION.txt)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_UNTRACKED_CHANGES:= $(shell git status --porcelain --untracked-files=no)
ifneq ($(GIT_UNTRACKED_CHANGES),)
	GIT_COMMIT := $(GIT_COMMIT)-dirty
endif
CTIMEVAR=-X $(PKG)/pkg/version.gitCommit=$(GIT_COMMIT) -X $(PKG)/pkg/version.version=$(VERSION)

CGO_ENABLED ?= 0
GO_LDFLAGS=-ldflags "-w $(CTIMEVAR)"
GO_LDFLAGS_STATIC=-ldflags "-w $(CTIMEVAR) -extldflags -static"
GO_BUILDTAGS := osusergo
GO_FLAGS ?= -tags '$(GO_BUILDTAGS)'

ifneq ($(wildcard go.mod),)  # exist go.mod
ifeq ($(CI),)  # $CI is empty
	GO_FLAGS+=-mod=vendor
endif
endif

GO_TEST ?= go test
GO_TEST_FUNC ?= .
GO_TEST_FLAGS ?=
GO_BENCH_FUNC ?= .
GO_BENCH_FLAGS ?= -benchmem

IMAGE_REGISTRY := gcr.io/container-image

# ----------------------------------------------------------------------------
# defines

GOPHER = ""
define target
@printf "$(GOPHER)  \\033[32m$(patsubst ,$@,$(1))\\033[0m\\n"
endef

# ----------------------------------------------------------------------------
# targets

## build and install

.PHONY: $(APP)
$(APP): $(wildcard *.go) $(wildcard */*.go) VERSION.txt
	$(call target)
	CGO_ENABLED=$(CGO_ENABLED) go build $(strip $(GO_FLAGS)) -o $(APP) $(CMD)

.PHONY: build
build: GO_FLAGS+=${GO_LDFLAGS}
build: $(APP)  ## Builds a dynamic executable or package.

.PHONY: static
static: GO_FLAGS+=${GO_LDFLAGS_STATIC}
static: GO_BUILDTAGS+=static
static: $(APP)  ## Builds a static executable or package.

.PHONY: install
install: GO_FLAGS+=${GO_LDFLAGS_STATIC}
install:  ## Installs the executable or package.
	$(call target)
	CGO_ENABLED=$(CGO_ENABLED) go install -a -v $(strip $(GO_FLAGS)) $(CMD)


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
coverage:  ## Take test coverage.
	$(call target)
	$(GO_TEST) -v -race $(strip $(GO_FLAGS)) -covermode=atomic -coverpkg=$(PKG)/... -coverprofile=coverage.out $(GO_PKGS)

$(GO_PATH)/bin/go-junit-report:
	@GO111MODULE=off go get -u github.com/jstemmer/go-junit-report

.PHONY: cmd/go-junit-report
cmd/go-junit-report: $(GO_PATH)/bin/go-junit-report  # go get 'go-junit-report' binary

.PHONY: coverage/junit
coverage/junit: cmd/go-junit-report  ## Take test coverage and output test results with junit syntax.
	$(call target)
	@mkdir -p test-results
	$(GO_TEST) -v -race $(strip $(GO_FLAGS)) -covermode=atomic -coverpkg=$(PKG)/... -coverprofile=coverage.out $(GO_PKGS) 2>&1 | tee /dev/stderr | go-junit-report -set-exit-code > test-results/report.xml


## lint

.PHONY: lint
lint: lint/fmt lint/govet lint/golangci-lint  ## Run all linters.

.PHONY: lint/fmt
lint/fmt:  ## Verifies all files have been `gofmt`ed.
	$(call target)
	@gofmt -s -l . | grep -v -e 'vendor' -e '.pb.go' -e '_*.go' | tee /dev/stderr

.PHONY: lint/govet
lint/govet:  ## Verifies `go vet` passes.
	$(call target)
	@go vet -all $(GO_PKGS) | tee /dev/stderr

$(GO_PATH)/bin/golangci-lint:
	@GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

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
	@GO111MODULE=on go mod init

.PHONY: mod/goget
mod/goget:  ## Update module and go.mod.
	$(call target)
	@GO111MODULE=on go get -u -m -v -x ./...

.PHONY: mod/tidy
mod/tidy:
	$(call target)
	@GO111MODULE=on go mod tidy -v

.PHONY: mod/vendor
mod/vendor:
	$(call target)
	@GO111MODULE=on go mod vendor -v

.PHONY: mod/graph
mod/graph:
	$(call target)
	@GO111MODULE=on go mod graph

.PHONY: mod/clean
mod/clean:
	$(call target)
	@$(RM) go.mod go.sum
	@$(RM) -r vendor

.PHONY: mod
mod: mod/clean mod/init mod/tidy mod/vendor  ## Updates the vendoring directory via go mod.
	@sed -i ':a;N;$$!ba;s|go 1\.12\n\n||g' go.mod

.PHONY: mod/install
mod/install: mod/tidy mod/vendor
	$(call target)
	@GO111MODULE=off go install -v $(GO_VENDOR_PKGS) || GO111MODULE=on go install -mod=vendor -v $(GO_VENDOR_PKGS)

.PHONY: mod/update
mod/update: mod/goget mod/tidy mod/vendor mod/install  ## Updates all vendor packages.


## dep

$(GO_PATH)/bin/dep:
	@go get -u github.com/golang/dep/cmd/dep

.PHONY: cmd/dep
cmd/dep: $(GO_PATH)/bin/dep

.PHONY: dep/init
dep/init: cmd/dep  ## Fetch vendor packages via dep ensure.
	$(call target)
	@dep init -v -no-examples

.PHONY: dep/ensure
dep/ensure: cmd/dep Gopkg.toml  ## Fetch vendor packages via dep ensure.
	$(call target)
	@dep ensure -v

.PHONY: dep/ensure/vendor-only
dep/ensure/vendor-only: cmd/dep Gopkg.toml Gopkg.lock  ## Fetch vendor packages via dep ensure.
	$(call target)
	@dep ensure -v -vendor-only

.PHONY: dep/install
dep/install: dep/ensure  ## Compiles vendor packages to object(.a) file.
	@go install -v  $(GO_VENDOR_PKGS)

.PHONY: dep/update
dep/update: cmd/dep  ## Updates the vendoring directory via dep
	$(call target)
	@dep ensure -v -update

.PHONY: dep/clean
dep/clean: cmd/dep  ## Cleans the dep files.
	$(call target)
	@$(RM) Gopkg.toml Gopkg.lock
	@$(RM) -r vendor

.PHONY: dep
dep: dep/init dep/ensure dep/install  ## Initialize dep vendor structures.


## miscellaneous

.PHONY: container/build
container/build: Dockerfile  ## Create the container image from the Dockerfile.
	docker image build --rm --force-rm --pull --progress=plain -t $(IMAGE_REGISTRY)/$(APP):$(VERSION:v%=%) .

.PHONY: container/push
container/push:  ## Push the container image to $IMAGE_REGISTRY.
	docker image push $(IMAGE_REGISTRY)/$(APP):$(VERSION:v%=%)


.PHONY: boilerplate/go/%
boilerplate/go/%: BOILERPLATE_PKG_DIR=$(shell printf $@ | cut -d'/' -f3- | rev | cut -d'/' -f2- | rev)
boilerplate/go/%: BOILERPLATE_PKG_NAME=$(if $(findstring $@,cmd),main,$(shell printf $@ | rev | cut -d/ -f2 | rev))
boilerplate/go/%: hack/boilerplate/boilerplate.go.txt  ## Create go file from boilerplate.go.txt
	@if [ ! -d ${BOILERPLATE_PKG_DIR} ]; then mkdir -p ${BOILERPLATE_PKG_DIR}; fi
	@cat hack/boilerplate/boilerplate.go.txt <(printf "package ${BOILERPLATE_PKG_NAME}\\n") > $*


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
