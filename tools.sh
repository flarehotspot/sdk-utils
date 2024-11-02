#!/usr/bin/env bash

echo "Installing CLI tools..."

tools=(
  github.com/cespare/reflex@v0.3.1
)

cd ./core

for tool in ${tools[@]}; do
  go install -buildvcs=false $tool
done
