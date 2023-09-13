#!/bin/sh

WORKDIR=$(pwd)
GO_VERSION=$(cat "${WORKDIR}/go-version")

# Get the first two numbers of the version
# GO_SHORT_VERSION="$(cut -d '.' -f 1 <<< "$GO_VERSION")"."$(cut -d '.' -f 2 <<< "$GO_VERSION")"
GO_SHORT_VERSION=$(echo "$GO_VERSION" | awk -F. '{print $1"."$2}')

GOWORK="go ${GO_SHORT_VERSION}
use (
    ./core
    ./main"

    if [ -d ./plugins ]; then

        for d in ./plugins/*;
        do
            PLUGIN="$(basename $d)"
            GOWORK="${GOWORK}
            ${d}"
        done
    fi

    GOWORK="${GOWORK}
)"

echo "$GOWORK" > "$WORKDIR/go.work"

echo "go.work file created."

wait

cd $WORKDIR
