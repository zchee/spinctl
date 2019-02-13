# ----------------------------------------------------------------------------
# global

.DEFAULT_GOAL := static
APP = spinctl
CMD_PREFIX = $(PKG)/cmd/
CMD = $(CMD_PREFIX)$(APP)

# ----------------------------------------------------------------------------
# target

KUBERNETES_VERSION ?= kubernetes-1.13.3
mod/k8s:
	$(call target)
	@GO111MODULE=on go get -u -m -v -x k8s.io/api@${KUBERNETES_VERSION} k8s.io/apimachinery@${KUBERNETES_VERSION} k8s.io/cli-runtime@${KUBERNETES_VERSION} k8s.io/client-go@${KUBERNETES_VERSION}

# ----------------------------------------------------------------------------
# overlays

.PHONY: mod/update
mod/update: mod/goget mod/k8s mod/tidy mod/vendor mod/install  ## Updates all vendor packages.

# ----------------------------------------------------------------------------
# include

include hack/make/go.mk
