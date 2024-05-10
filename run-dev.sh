#!/bin/sh

GO_TAGS=${GO_TAGS:-"dev mono"}
FLARE_CLI=./core/internal/cli/flare-internal.go

go run $FLARE_CLI make-mono && \
    go run -tags="$GO_TAGS" ./main/main.go
