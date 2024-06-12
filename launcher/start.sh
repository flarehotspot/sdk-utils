#!/bin/bash

go run cmd/main.go && \
    cd app && \
    ./bin/flare server
