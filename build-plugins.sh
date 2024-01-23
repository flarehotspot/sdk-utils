#!/bin/sh

# This script builds all plugins in the plugins directory (plugin.so) and copies them to the vendor directory.

WORKDIR=$(pwd)

echo "Using go from $(which go)..."

for d in vendor/*; do
    PLUGIN=$(basename "$d")
    PLUGINDIR="$WORKDIR/plugins/$PLUGIN"
    VENDORDIR="$WORKDIR/vendor/$PLUGIN"

    if [ -d $PLUGINDIR ]; then
        if [ -e $VENDORDIR ]; then
            rm -rf $VENDORDIR
        fi

        echo "Building plugin $d..." && \
            cd $PLUGINDIR && go build -buildmode=plugin -ldflags="-s -w" -trimpath -o plugin.so ./main.go && \
            cp -r $PLUGINDIR $VENDORDIR
    fi
done

wait

cd $WORKDIR
