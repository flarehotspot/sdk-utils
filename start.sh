#!/usr/bin/env bash

BUILD_TAGS="dev"
BUILD_CORE_MAIN="./core/cmd/build-core"
BUILD_CLI_MAIN="./core/cmd/build-cli"
BUILD_TEMPLATES="./core/cmd/build-templates"
FLARE_BIN="./bin/flare"

sqlc generate && \
    go run -tags="${BUILD_TAGS}" $BUILD_CLI_MAIN && \
    sh -c "$FLARE_BIN fix-workspace" && \
    sh -c "$FLARE_BIN build-templates" && \
    go run -tags="${BUILD_TAGS}" $BUILD_CORE_MAIN && \
    sh -c "$FLARE_BIN server"
