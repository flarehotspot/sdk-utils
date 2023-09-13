#!/bin/sh

DOCKER_IMAGE="devkit:latest"
TMP_CONTAINER="devkit-tmp"
CORE_SO="/root/core.so"
OUTFILE="devkit/core/core.so"
RELEASE_DIR="devkit-release"
DEVKIT_FILES=(
    ./main
    ./core/go.mod
    ./core/go.sum
    ./core/sdk
    ./core/resources
    ./core/package.yml
    ./run.sh
    ./build-main.sh
    ./build-plugins.sh
    ./link-resources.sh
    ./go-work.sh
    ./go-version
    ./install-go.sh
    ./package.json
    ./package-lock.json
    ./.files
)

function copy_main_filess() {
    mkdir -p $RELEASE_DIR/core

    for file in "${DEVKIT_FILES[@]}"; do
        echo "Copying $file"
        cp -r $file $RELEASE_DIR/$file
    done
}

function default_configs() {
    secret=$(openssl rand -hex 16)
    mkdir -p $RELEASE_DIR/config
    cp -r ./config/.defaults/ $RELEASE_DIR/config/.defaults

    cat > $RELEASE_DIR/config/application.yml<<EOF
---
secret: $secret
lang: en
EOF

    echo "Created config/application.yml:" && \
        echo $(cat $RELEASE_DIR/config/application.yml)
}

function copy_devkit_files() {
    echo "Copying devkit files..."
    cp -r ./devkit/* $RELEASE_DIR/
}


function prepare() {
    rm -rf $RELEASE_DIR && \
        mkdir -p $RELEASE_DIR/plugins && \
        docker rm -f "$TMP_CONTAINER" || true
}

prepare && \
    docker build -t "$DOCKER_IMAGE" . && \
    docker cp $(docker create --name ${TMP_CONTAINER} ${DOCKER_IMAGE}):${CORE_SO} $OUTFILE && \
    docker rm "$TMP_CONTAINER" && \
    default_configs && \
    copy_main_filess && \
    copy_devkit_files
