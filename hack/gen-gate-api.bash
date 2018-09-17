#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# Usage:
#
# $ ./hack/gen-gate-api.bash 'http://localhost:8084' '1.9.3' 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx'
#
# NOTE: The 'SESSION' cookie stored browser cookie data

ENDPOINT="$1"
SPINNAKER_VERSION="$2"
SESSION="$3"
GO_PATH=${GOPATH:-$HOME/go}
PKG_ROOT="$GO_PATH/src/github.com/zchee/spinctl"

rm -rf "$PKG_ROOT/api/gate"
swagger-codegen generate --auth "cookie: SESSION=$SESSION" -i "$ENDPOINT/v2/api-docs" -l go -D 'packageName=gate' -D "packageVersion=$SPINNAKER_VERSION" -o "$PKG_ROOT/api/gate"

curl -v -H "cookie: SESSION=$SESSION" -o "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION-raw.json" "$ENDPOINT/v2/api-docs"
if [[ -f "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION.json" ]]; then
  rm -f "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION.json"
fi
jq ' ' < "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION-raw.json" > "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION.json"
rm -f "$PKG_ROOT/api/swagger-spec/gate-$SPINNAKER_VERSION-raw.json"

goimports -w api/gate
