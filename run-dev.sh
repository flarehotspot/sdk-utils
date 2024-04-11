#!/bin/sh

FLARE_CLI=./core/internal/cli/flare-internal.go

go run $FLARE_CLI make-mono && \
    go run -tags="dev mono" ./main/main.go
