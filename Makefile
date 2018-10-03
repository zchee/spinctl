# ----------------------------------------------------------------------------
# global environment variables
 
GO_PATH := $(shell go env GOPATH)

APP = spinctl
PKG_ROOT = github.com/zchee/spinctl

GO_PKGS := $(shell go list ./... | grep -v -e 'api/gate')
GO_PGKS_ABS := $(shell go list -f '${GO_PATH}/src/{{.ImportPath}}' ./... | grep -v -e 'api/gate')

# ----------------------------------------------------------------------------
# targets

all: build

$(APP): ${GO_PGKS_ABS}
	go build -v -o $@ ${PKG_ROOT}/cmd/${APP}

build: ${APP}  ## build spinctl binary


LINTERS = deadcode dupl errcheck goconst gocyclo golint gosec ineffassign interfacer maligned megacheck structcheck unconvert varcheck 
lint: golangci-lint  ## Run all linters

$(GO_PATH)/bin/golangci-lint:
	@go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint

golangci-lint: $(GO_PATH)/bin/golangci-lint
	golangci-lint run --no-config --issues-exit-code=0 --deadline=30m --disable-all $(foreach tool,$(LINTERS),--enable=$(tool)) $(GO_PGKS_ABS)

# ----------------------------------------------------------------------------
# help

.PHONY: help
help:  ## show make target help
	@perl -nle 'BEGIN {printf "Usage:\n  make \033[33m<target>\033[0m\n\nTargets:\n"} printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 if /^([a-zA-Z\/_-].+):\s+## (.*)/' ${MAKEFILE_LIST}
