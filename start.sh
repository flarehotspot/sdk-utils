#!/bin/env bash

BUILD_TAGS="dev mono"
FLARECLI="./core/internal/cli/flare-internal.go"

go run -tags="${BUILD_TAGS}" $FLARECLI make-mono && \
    go run -tags="${BUILD_TAGS}" $FLARECLI server


