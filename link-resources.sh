#!/bin/sh

WORKDIR=$(pwd)

for d in vendor/*;
do
    PLUGIN=$(basename "$d")

    if [[ -e "$WORKDIR/plugins/$PLUGIN" ]]; then
        VENDOR_RESOURCES="$WORKDIR/vendor/$PLUGIN/resources"
        PLUGIN_RESOURCES="$WORKDIR/plugins/$PLUGIN/resources"

        echo "Linking files $VENDOR_RESOURCES -> $PLUGIN_RESOURCES"

        rm -rf "$VENDOR_RESOURCES"
        ln -s $PLUGIN_RESOURCES $VENDOR_RESOURCES
    fi
done

wait

cd $WORKDIR
