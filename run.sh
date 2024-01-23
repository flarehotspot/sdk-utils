#!/bin/sh

WORKDIR=$(pwd)

# Clean up
rm -rf .tmp .cache/views public
rm -rf ./vendor && mkdir ./vendor
find . -name "*.app" -type f -delete

# Build .so files and run
./go-work.sh
./build-main.sh
./build-plugins.sh
./main/main.app
