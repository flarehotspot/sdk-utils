#!/bin/sh

echo "Installing CLI tools..."
go install github.com/cespare/reflex@v0.3.1 && \
go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0
