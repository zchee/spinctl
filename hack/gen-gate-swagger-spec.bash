#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'
set -x

# Usage:
#
# $ ./hack/gen-gate-swagger-spec.bash 'http://localhost:8084' '1.9.3' 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx'
#
# NOTE: The 'SESSION' cookie stored browser cookie data

ENDPOINT="$1"
SPINNAKER_VERSION="$2"
SESSION="$3"
GO_PATH=${GOPATH:-$HOME/go}
PKG_ROOT="$GO_PATH/src/github.com/zchee/spinctl"

curl -sv -H "cookie: SESSION=$SESSION" -o "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION-raw.json" "$ENDPOINT/v2/api-docs"
if [[ -f "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION.json" ]]; then
  rm -f "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION.json"
fi
jq ' ' < "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION-raw.json" > "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION.json"
rm -f "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION-raw.json"

git add api/swagger-spec/"gate-$SPINNAKER_VERSION.json"
git commit -m "api/swagger-spec: add gate swagger spec from spinnaker $SPINNAKER_VERSION"
