#!/usr/bin/env bash

GO="./go/bin/go"
BUILD_TAGS="dev"
CORE_MAIN="./core/cmd/build-core/main.go"
CLI_MAIN="./core/cmd/build-cli/main.go"
FLARE_BIN="./bin/flare"

$GO run -tags="${BUILD_TAGS}" $CORE_MAIN && \
    $GO run -tags="${BUILD_TAGS}" $CLI_MAIN && \
    sh -c "$FLARE_BIN fix-workspace" && \
    sh -c "$FLARE_BIN server"
