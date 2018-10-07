# ----------------------------------------------------------------------------
# global environment variables
 
APP = spinctl
PKG_ROOT = github.com/zchee/spinctl

GO_PATH := $(shell go env GOPATH)
GO_ALL_PGKS := $(shell go list -f '$(GO_PATH)/src/{{.ImportPath}}' ./...)
GO_PKGS := $(shell go list ./... | grep -v -e 'api/gate')
GO_PGKS_ABS := $(shell go list -f '$(GO_PATH)/src/{{.ImportPath}}' ./... | grep -v -e 'api/gate')
GO_TEST ?= go test
GO_TEST_FUNC ?= .

# ----------------------------------------------------------------------------
# defines

define target
	@printf "+ \\033[32m$(shell printf $@ | cut -d '/' -f2)\\033[0m\\n"
endef

# ----------------------------------------------------------------------------
# targets

$(APP): $(GO_ALL_PGKS)
	$(call target)
	go build -v -o $@ $(PKG_ROOT)/cmd/$(APP)

build: $(APP)  ## build spinctl binary


.PHONY: test
test:  ## Run the package test
	$(call target)
	$(GO_TEST) -v -race -run=$(GO_TEST_FUNC) $(GO_PKGS)

.PHONY: coverage
coverage:  ## Take test coverage
	$(call target)
	$(GO_TEST) -v -race -covermode=atomic -coverpkg=./... -coverprofile=coverage.out $(GO_PKGS)


lint: lint/golangci-lint  ## Run all linters

$(GO_PATH)/bin/golangci-lint:
	@go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: golangci-lint
LINTERS = deadcode dupl errcheck goconst gocyclo golint gosec ineffassign interfacer maligned megacheck structcheck unconvert varcheck 
lint/golangci-lint: $(GO_PATH)/bin/golangci-lint  ## Run golangci-lint
	$(call target)
	golangci-lint run --no-config --issues-exit-code=0 --deadline=30m --disable-all $(foreach tool,$(LINTERS),--enable=$(tool)) $(GO_PGKS_ABS)


$(GO_PATH)/bin/dep:
	@go get -u -v github.com/golang/dep/cmd/dep

.PHONY: dep/ensure
dep/ensure: $(GO_PATH)/bin/dep  ## Fetch vendor packages via dep ensure
	$(call target)
	dep ensure -v -vendor-only

.PHONY: dep/update
dep/update: $(GO_PATH)/bin/dep  ## Update vendor packages
	$(call target)
	dep ensure -v -update


.PHONY: clean
clean:  ## clean workspace
	$(call target)
	@rm -f ./spinctl *.out *.prof


.PHONY: help
help:  ## show make target help
	@perl -nle 'BEGIN {printf "Usage:\n  make \033[33m<target>\033[0m\n\nTargets:\n"} printf "  \033[32m%-30s\033[0m %s\n", $$1, $$2 if /^([a-zA-Z\/_-].+)+:.*?\s+## (.*)/' ${MAKEFILE_LIST}
