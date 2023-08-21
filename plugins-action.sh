#!/bin/bash

ACTION=$1
WORKING_DIR=$(pwd)

function linkresources() {
  ln -s $WORKING_DIR/plugins/$1/resources $WORKING_DIR/vendor/$1/resources
}

for d in plugins/*; do \
  echo "cd $WORKING_DIR/$d && $ACTION";\
  cd $WORKING_DIR/$d && $ACTION &\
done

wait
