#!/bin/sh

WORKING_DIR=$(pwd)

for d in plugins/*; do \
    echo "cd $WORKING_DIR/$d && make plugin"; \
    cd $WORKING_DIR/$d && make plugin &\
done

wait
