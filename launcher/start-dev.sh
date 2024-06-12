#!/bin/bash

go run --tags=dev cmd/main.go && \
    cd app && \
    ./bin/flare server
