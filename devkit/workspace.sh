#!/bin/sh

WORKING_DIR=$(pwd)

GOWORK="go 1.19
use (
  ./main
  ./sdk"

for d in plugins/*; do \
  GOWORK="$GOWORK
  ./$d"
done
  GOWORK="$GOWORK
)"

echo "$GOWORK" > go.work

wait
