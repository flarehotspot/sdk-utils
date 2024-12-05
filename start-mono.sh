#!/usr/bin/env bash

BUILD_TAGS="dev mono"
CREATE_MONO="./core/cmd/make-mono/main.go"
MONO_SERVER="./core/cmd/mono-server/main.go"

go run -tags="${BUILD_TAGS}" $CREATE_MONO && \
    go run -tags="${BUILD_TAGS}" $MONO_SERVER
