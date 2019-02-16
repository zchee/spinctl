# ----------------------------------------------------------------------------
# global

.DEFAULT_GOAL := static
APP = spinctl
CMD_PREFIX = $(PKG)/cmd/
CMD = $(CMD_PREFIX)$(APP)

GATE_FORWARDED_HOST ?= 'http://localhost:8084'
GATE_VERSION ?=
GATE_SESSION ?=

# ----------------------------------------------------------------------------
# defines

define error
@printf "$(GOPHER)  : \x1b[1;31m[ERROR]\x1b[0m %s\n\\" "$1"
endef

# ----------------------------------------------------------------------------
# target

KUBERNETES_VERSION ?= kubernetes-1.13.3
mod/k8s:  ## Override k8s components version
	$(call target)
	@GO111MODULE=on go get -u -m -v -x k8s.io/api@${KUBERNETES_VERSION} k8s.io/apimachinery@${KUBERNETES_VERSION} k8s.io/cli-runtime@${KUBERNETES_VERSION} k8s.io/client-go@${KUBERNETES_VERSION}

.PHONY: hack/gen-gate-api
hack/gen-gate-api:
ifeq (${GATE_VERSION},)
	$(error required GATE_VERSION env)
endif
ifeq (${GATE_SESSION},)
	$(error required GATE_SESSION cookie)
endif
	@bash hack/gen-gate-api.bash ${GATE_FORWARDED_HOST} '${GATE_VERSION}' '${GATE_SESSION}'


# ----------------------------------------------------------------------------
# overlays

.PHONY: mod/update
mod/update: mod/goget mod/k8s mod/tidy mod/vendor mod/install  ## Updates all vendor packages.

# ----------------------------------------------------------------------------
# include

include hack/make/go.mk
