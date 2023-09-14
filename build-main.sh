#!/bin/sh

WORKDIR=$(pwd)

cd main && go build -ldflags="-s -w" -tags="dev" -trimpath -o main.app main.go

cd $WORKDIR
