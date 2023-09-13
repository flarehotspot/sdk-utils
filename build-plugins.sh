#!/bin/sh

WORKDIR=$(pwd)
PATH="$WORKDIR/go/bin:$PATH"

echo "Using go from $(which go)..."

for d in plugins/*; do
    PLUGIN=$(basename "$d")
    echo "Building plugin $d..." && \
        cd $WORKDIR/$d && go build -buildmode=plugin -ldflags="-s -w" -trimpath -o plugin.so ./main.go && \
        cp -r $WORKDIR/$d $WORKDIR/vendor/$PLUGIN
done

wait

cd $WORKDIR
