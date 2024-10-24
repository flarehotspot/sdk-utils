#!/usr/bin/env bash

BUILD_TAGS="dev"
CORE_MAIN="./core/cmd/build-core/main.go"
CLI_MAIN="./core/cmd/build-cli/main.go"
FLARE_BIN="./bin/flare"

go run -tags="${BUILD_TAGS}" $CORE_MAIN && \
    go run -tags="${BUILD_TAGS}" $CLI_MAIN && \
    sh -c "$FLARE_BIN fix-workspace" && \
    sh -c "$FLARE_BIN server"
