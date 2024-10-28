#!/usr/bin/env bash

GO="./go/bin/go"
BUILD_TAGS="dev"
BUILD_CORE="./core/cmd/build-core/main.go"
BUILD_CLI="./core/cmd/build-cli/main.go"
FLARE_BIN="./bin/flare"

$GO run -tags="${BUILD_TAGS}" $BUILD_CORE && \
    $GO run -tags="${BUILD_TAGS}" $BUILD_CLI && \
    sh -c "$FLARE_BIN server"
