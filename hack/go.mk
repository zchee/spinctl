# ----------------------------------------------------------------------------
# global
 
GO_PATH = $(shell go env GOPATH)
PKG = $(subst $(GO_PATH)/src/,,$(CURDIR))
GO_PKGS := $(shell go list ./... | grep -v -e '.pb.go')
GO_PKGS_ABS := $(shell go list -f '$(GO_PATH)/src/{{.ImportPath}}' ./... | grep -v -e '.pb.go')

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
ifeq ($(GIT_COMMIT),)
    GIT_COMMIT := ${GITHUB_SHA}
endif
CTIMEVAR=-X $(PKG)/pkg/version.gitCommit=$(GITCOMMIT) -X $(PKG)/pkg/version.version=$(VERSION)
GO_LDFLAGS=-ldflags "-w $(CTIMEVAR)"
GO_LDFLAGS_STATIC=-ldflags "-w $(CTIMEVAR) -extldflags -static"

VET_LINTERS := asmdecl assign atomic bools buildtag cgocall copylocks httpresponse loopclosure lostcancel nilfunc nilness pkgfact shift stdmethods structtag tests unreachable unsafeptr  # composites -composites.whitelist '' findcall -findcall.name '' printf -printf.funcs '' unusedresult -unusedresult.funcs '' -unusedresult.stringmethods ''
GOLANGCI_LINTERS := deadcode dupl errcheck goconst gocyclo golint gosec ineffassign interfacer maligned megacheck structcheck unconvert varcheck 

IMAGE_REGISTRY := quay.io/zchee

# ----------------------------------------------------------------------------
# defines

define target
@printf "+ \\033[32m$@\\033[0m\\n"
endef

# ----------------------------------------------------------------------------
# targets

$(APP): $(wildcard *.go) $(wildcard */**/*.go) VERSION.txt
	$(call target)
	CGO_ENABLED=$(CGO_ENABLED) go build -tags "$(BUILDTAGS)" -v $(GO_LDFLAGS) -o $(APP) $(PKG)/cmd/$(APP)

build: $(APP)  ## Builds a dynamic executable or package.

static: CGO_ENABLED=0
static: GO_LDFLAGS=${GO_LDFLAGS_STATIC}
static: BUILDTAGS+=static
static: $(APP)  ## Builds a static executable or package.

.PHONY: install
install: GO_LDFLAGS=${GO_LDFLAGS_STATIC}
install:  ## Installs the executable or package.
	$(call target)
	go install -a -tags "$(BUILDTAGS)" ${GO_LDFLAGS} $(PKG)/cmd/$(APP)


.PHONY: test
test: BUILDTAGS+=cgo
test:  ## Run the package test.
	$(call target)
	$(GO_TEST) -v -race -tags "$(BUILDTAGS)" -run=$(GO_TEST_FUNC) $(GO_PKGS)

.PHONY: bench
bench: BUILDTAGS+=cgo
bench:  ## Take a package benchmark.
	$(call target)
	$(GO_TEST) -v -race -tags "$(BUILDTAGS)" -run='^$$' -bench=$(GO_BENCH_FUNC) -benchmem $(GO_PKGS)

.PHONY: coverage
coverage:  ## Take test coverage.
	$(call target)
	$(GO_TEST) -v -race -covermode=atomic -coverpkg=$(PKG)/... -coverprofile=coverage.out $(GO_PKGS)


lint: lint/fmt lint/govet lint/golint lint/golangci-lint  ## Run all linters.

.PHONY: lint/fmt
lint/fmt:  ## Verifies all files have been `gofmt`ed.
	$(call target)
	@gofmt -s -l . | grep -v '.pb.go' | tee /dev/stderr

.PHONY: lint/govet
lint/govet:  ## Verifies `go vet` passes.
	$(call target)
	@go vet $(shell $(GO) list ./... | grep -v vendor) | grep -v '.pb.go:' | tee /dev/stderr

$(GO_PATH)/bin/golint:
	@go get -u golang.org/x/lint/golint

.PHONY: lint/golint
lint/golint: $(GO_PATH)/bin/golint  ## Verifies `golint` passes.
	$(call target)
	@golint ./... | grep -v '.pb.go' | grep -v vendor | tee /dev/stderr

$(GO_PATH)/bin/vet:
	@go get -u golang.org/x/tools/go/analysis/cmd/vet golang.org/x/tools/go/analysis/passes/...

lint/vet:  ## Run vet
	$(call target)
	@vet $(foreach linter,$(VET_LINTERS),-$(linter).enable) $(GO_PKGS)

$(GO_PATH)/bin/golangci-lint:
	@go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: golangci-lint
lint/golangci-lint: $(GO_PATH)/bin/golangci-lint  ## Run golangci-lint.
	$(call target)
	@golangci-lint run --no-config --issues-exit-code=0 $(foreach pat,$(shell cat .errcheckignore),--exclude '$(pat)') --deadline=30m --disable-all $(foreach tool,$(GOLANGCI_LINTERS),--enable=$(tool)) $(GO_PKGS_ABS)


$(GO_PATH)/bin/dep:
	@go get -u github.com/golang/dep/cmd/dep

.PHONY: dep/init
dep/init: $(GO_PATH)/bin/dep  ## Fetch vendor packages via dep ensure.
	$(call target)
	@dep init -v -no-examples

.PHONY: dep/ensure
dep/ensure: $(GO_PATH)/bin/dep  ## Fetch vendor packages via dep ensure.
	$(call target)
	@dep ensure -v -vendor-only

.PHONY: dep/update
dep/update: $(GO_PATH)/bin/dep dep/clean  ## Updates the vendoring directory.
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
