#!/bin/sh

WORKDIR=$(pwd)

# Clean up
rm -rf .tmp .cache/views public
rm -rf ./vendor && mkdir ./vendor
find . -name "*.app" -type f -delete

# Build .so files and run
node ./make-go.work.js
node ./build-main.js
node ./build-plugins.js
node ./run.js
