#!/bin/bash

WORKING_DIR=$(pwd)

for d in vendor/*;
do
    echo "\$d is: $d"
    echo "basename: $(basename $d)"
    PLUGIN="$(basename $d)"
    VENDOR_RESOURCES="$WORKING_DIR/vendor/$PLUGIN/resources"
    PLUGIN_RESOURCES="$WORKING_DIR/plugins/$PLUGIN/resources"

    echo "Linking files $VENDOR_RESOURCES -> $PLUGIN_RESOURCES"

    rm -rf "$VENDOR_RESOURCES"
    ln -s $PLUGIN_RESOURCES $VENDOR_RESOURCES
done

wait
