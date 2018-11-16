# ----------------------------------------------------------------------------
# global
 
GO_ROOT = $(shell go env GOROOT)
GO_PATH = $(shell go env GOPATH)
GO_GOOS = $(shell go env GOOS)
GO_GOARCH = $(shell go env GOARCH)
PKG = $(subst $(GO_PATH)/src/,,$(CURDIR))
GO_PKGS := $(shell go list ./... | grep -v -e '.pb.go' -e 'api/gate' -e 'api/mock')
GO_PKGS_ABS := $(shell go list -f '$(GO_PATH)/src/{{.ImportPath}}' ./... | grep -v -e '.pb.go' -e 'api/gate' -e 'api/mock')

GO_TEST ?= go test
GO_TEST_FUNC ?= .
GO_BENCH_FUNC ?= .

CGO_ENABLED := 1
VERSION := $(shell cat VERSION.txt)

GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_UNTRACKED_CHANGES:= $(shell git status --porcelain --untracked-files=no)
ifneq ($(GIT_UNTRACKED_CHANGES),)
	GIT_COMMIT := $(GIT_COMMIT)-dirty
endif
CTIMEVAR=-X $(PKG)/pkg/version.gitCommit=$(GIT_COMMIT) -X $(PKG)/pkg/version.version=$(VERSION)
GO_LDFLAGS=-ldflags "-w $(CTIMEVAR)"
GO_LDFLAGS_STATIC=-ldflags "-w $(CTIMEVAR) -extldflags -static"
GO_BUILDTAGS := osusergo

VETLITE_LINTERS ?= asmdecl assign atomic bools buildtag cgocall copylocks httpresponse loopclosure lostcancel nilfunc pkgfact shift stdmethods structtag tests unmarshal unreachable unsafeptr ## composites composites.whitelist printf printf.funcs unusedresult unusedresult.funcs unusedresult.stringmethods
GOLANGCI_ENABLE_LINTERS ?= deadcode,depguard,dupl,errcheck,goconst,gocritic,gocyclo,gofmt,goimports,golint,gosec,gosimple,govet,ineffassign,interfacer,maligned,misspell,nakedret,prealloc,scopelint,staticcheck,structcheck,unconvert,unparam,unused,varcheck
GOLANGCI_DISABLE_LINTERS ?= gochecknoglobals,gochecknoinits,lll,megacheck,typecheck
GOLANGCI_EXCLUDE ?=
ifneq ($(wildcard '.errcheckignore'),)
	GOLANGCI_EXCLUDE = $(foreach pat,$(shell cat .errcheckignore),--exclude '$(pat)')
endif

IMAGE_REGISTRY := quay.io/zchee

# ----------------------------------------------------------------------------
# defines

define target
@printf "+ \\033[32m$@\\033[0m\\n"
endef

# ----------------------------------------------------------------------------
# targets

debug:
	@echo PKG
	@echo $(PKG)
	@echo GO_PKGS
	@echo $(GO_PKGS)
	@echo GO_PKGS_ABS
	@echo $(GO_PKGS_ABS)

$(APP): $(wildcard *.go) $(wildcard */**/*.go) VERSION.txt
	$(call target)
	CGO_ENABLED=$(CGO_ENABLED) go build -tags "$(GO_BUILDTAGS)" -v $(GO_LDFLAGS) -o $(APP) $(PKG)/cmd/$(APP)

build: $(APP)  ## Builds a dynamic executable or package.

static: CGO_ENABLED=0
static: GO_LDFLAGS=${GO_LDFLAGS_STATIC}
static: GO_BUILDTAGS+=static
static: $(APP)  ## Builds a static executable or package.

.PHONY: install
install: GO_LDFLAGS=${GO_LDFLAGS_STATIC}
install:  ## Installs the executable or package.
	$(call target)
	go install -a -tags "$(BUILDTAGS)" ${GO_LDFLAGS} $(PKG)/cmd/$(APP)


.PHONY: test
test:  ## Run the package test.
	$(call target)
	$(GO_TEST) -v -race -tags "$(BUILDTAGS)" -run=$(GO_TEST_FUNC) $(GO_PKGS)

.PHONY: bench
bench:  ## Take a package benchmark.
	$(call target)
	$(GO_TEST) -v -race -tags "$(BUILDTAGS)" -run='^$$' -bench=$(GO_BENCH_FUNC) -benchmem $(GO_PKGS)

.PHONY: coverage
coverage:  ## Take test coverage.
	$(call target)
	$(GO_TEST) -v -race -covermode=atomic -coverpkg=$(PKG)/... -coverprofile=coverage.out $(GO_PKGS)


lint: lint/fmt lint/govet lint/golint lint/vet-lite lint/golangci-lint  ## Run all linters.

.PHONY: lint/fmt
lint/fmt:  ## Verifies all files have been `gofmt`ed.
	$(call target)
	@gofmt -s -l . | grep -v -e 'vendor' -e '.pb.go' | tee /dev/stderr

.PHONY: lint/govet
lint/govet:  ## Verifies `go vet` passes.
	$(call target)
	@go vet -all $(GO_PKGS) | tee /dev/stderr

$(GO_PATH)/bin/golint:
	@go get -u -v golang.org/x/lint/golint

.PHONY: cmd/golint
cmd/golint: $(GO_PATH)/bin/golint

.PHONY: lint/golint
lint/golint: cmd/golint  ## Verifies `golint` passes.
	$(call target)
	@golint -min_confidence=0.3 -set_exit_status $(GO_PKGS)

$(GO_ROOT)/pkg/tool/$(GO_GOOS)_$(GO_GOARCH)/vet-lite:
	@go install -v $(GO_ROOT)/src/cmd/vendor/golang.org/x/tools/go/analysis/cmd/vet-lite

.PHONY: cmd/vet-lite
cmd/vet-lite: $(GO_ROOT)/pkg/tool/$(GO_GOOS)_$(GO_GOARCH)/vet-lite

lint/vet-lite: cmd/vet-lite  ## Run vet
	$(call target)
	@go vet -vettool=$(GO_ROOT)/pkg/tool/darwin_amd64/vet-lite $(foreach linter,$(VETLITE_LINTERS),-$(linter)) $(GO_PKGS)

$(GO_PATH)/bin/golangci-lint:
	@go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: cmd/golangci-lint
cmd/golangci-lint: $(GO_PATH)/bin/golangci-lint

.PHONY: golangci-lint
lint/golangci-lint: cmd/golangci-lint  ## Run golangci-lint.
	$(call target)
	@golangci-lint run --no-config --fast --deadline=30m $(strip $(GOLANGCI_EXCLUDE)) --skip-files='.*\.pb\.go' --skip-dirs='(api\/gate|api\/mock)' --issues-exit-code=0 --enable='$(GOLANGCI_ENABLE_LINTERS)' --disable='$(GOLANGCI_DISABLE_LINTERS)' ./...


$(GO_PATH)/bin/dep:
	@go get -u github.com/golang/dep/cmd/dep

.PHONY: cmd/dep
cmd/dep: $(GO_PATH)/bin/dep

.PHONY: dep/init
dep/init: cmd/dep  ## Fetch vendor packages via dep ensure.
	$(call target)
	@dep init -v -no-examples

.PHONY: dep/ensure
dep/ensure: cmd/dep  ## Fetch vendor packages via dep ensure.
	$(call target)
	@dep ensure -v -vendor-only

.PHONY: dep/update
dep/update: cmd/dep dep/clean  ## Updates the vendoring directory.
	$(call target)
	@dep ensure -v -update

.PHONY: dep/clean
dep/clean:
	$(call target)
	@$(RM) -r vendor


.PHONY: container/build
container/build: Dockerfile  ## Create the container image from the Dockerfile.
	docker image build --rm --force-rm --pull --progress=plain -t $(IMAGE_REGISTRY)/$(APP):$(VERSION:v%=%) .

.PHONY: container/push
container/push:  ## Push the container image to $IMAGE_REGISTRY.
	docker image push $(IMAGE_REGISTRY)/$(APP):$(VERSION:v%=%)


.PHONY: AUTHORS
AUTHORS:
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@echo "$(shell git log --format='\n%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf)" >> $@


.PHONY: clean
clean:  ## Cleanup any build binaries or packages.
	$(call target)
	$(RM) $(APP) *.out *.prof


.PHONY: help
help:  ## Show make target help.
	@perl -nle 'BEGIN {printf "Usage:\n  make \033[33m<target>\033[0m\n\nTargets:\n"} printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 if /^([a-zA-Z\/_-].+)+:.*?\s+## (.*)/' ${MAKEFILE_LIST}
