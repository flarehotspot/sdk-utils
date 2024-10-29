#!/usr/bin/env bash

echo "Installing CLI tools..." && \
    go install -buildvcs=false ./sdk/libs/reflex-0.3.1 && \
    go install -buildvcs=false ./sdk/libs/sqlc-1.26.0/cmd/sqlc && \
    go install -buildvcs=false ./sdk/libs/templ-0.2.778/cmd/templ
