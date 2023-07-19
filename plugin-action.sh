#!/bin/bash

ACTION=$1
WORKING_DIR=$(pwd)

for d in plugins/*; do \
  echo "cd $WORKING_DIR/$d && $ACTION";\
  cd $WORKING_DIR/$d && $ACTION;\
done
