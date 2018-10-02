# ----------------------------------------------------------------------------
# global environment variables
 
GO_PATH := $(shell go env GOPATH)

APP = spinctl
PKG_ROOT = github.com/zchee/spinctl

GO_PKGS := $(shell go list ./...)
GO_PGKS_ABS := $(shell go list -f '${GO_PATH}/src/{{.ImportPath}}' ./...)

# ----------------------------------------------------------------------------
# targets

all: build

$(APP): ${GO_PKGS}
	go build -v -x -o $@ ${PKG_ROOT}/cmd/${APP}

build: ${APP}  ## build spinctl binary

# ----------------------------------------------------------------------------
# help

.PHONY: help
help:  ## show make target help
	@perl -nle 'BEGIN {printf "Usage:\n  make \033[33m<target>\033[0m\n\nTargets:\n"} printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 if /^([a-zA-Z\/_-].+):\s+## (.*)/' ${MAKEFILE_LIST}
