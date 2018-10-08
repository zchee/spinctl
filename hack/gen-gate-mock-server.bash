#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Usage:
#
# $ ./hack/gen-gate-api.bash '1.9.3'
#
# NOTE: The 'SESSION' cookie stored browser cookie data

SPINNAKER_VERSION="$1"
GO_PATH=${GOPATH:-$HOME/go}
PKG_ROOT="$GO_PATH/src/github.com/zchee/spinctl"

rm -rf "$PKG_ROOT/api/mock/gate"
swagger-codegen generate -i "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION.json" -l go-server -D 'packageName=gatemock' -D "packageVersion=$SPINNAKER_VERSION" -o "$PKG_ROOT/api/mock/gate"

rm -f "$PKG_ROOT/api/mock/gate/main.go"
mv "$PKG_ROOT/api/mock/gate/go/routers.go" "$PKG_ROOT/api/mock/gate/go/_routers.go"
goimports -w api/mock/gate

git add api/mock/gate
git commit -m "api/mock/gate: regenerate gate API mock server from spinnaker $SPINNAKER_VERSION"
