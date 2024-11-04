#!/usr/bin/env bash

BUILD_TAGS="dev"
BUILD_CORE_MAIN="./core/cmd/build-core/main.go"
BUILD_CLI_MAIN="./core/cmd/build-cli/main.go"
FLARE_BIN="./bin/flare"

go run -tags="${BUILD_TAGS}" $BUILD_CLI_MAIN && \
    sh -c "$FLARE_BIN fix-workspace" && \
    go run -tags="${BUILD_TAGS}" $BUILD_CORE_MAIN && \
    sh -c "$FLARE_BIN server"
