#!/bin/sh

WORKDIR=$(pwd)
PATH="$WORKDIR/go/bin:$PATH"

cd main && go build -ldflags="-s -w" -tags="dev" -trimpath -o main.app main.go

cd $WORKDIR
