#!/usr/bin/env bash

BUILD_TAGS="dev mono"
CREATE_MONO="./core/cmd/make-mono/main.go"
DEBUG_SERVER="./core/cmd/debug-server/main.go"

go run -tags="${BUILD_TAGS}" $CREATE_MONO && \
    go run -tags="${BUILD_TAGS}" $DEBUG_SERVER


