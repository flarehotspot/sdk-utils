#!/usr/bin/env bash

PACKAGES=$(cat ./packages.txt)

go mod tidy

for package in $PACKAGES; do
    echo "installing package: $package"
    go install $package
done
