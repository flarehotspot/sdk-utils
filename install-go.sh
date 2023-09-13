#!/bin/sh

WORKDIR=$(pwd)
CACHE_PATH="${WORKDIR}/.cache"

GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)
GO_VERSION=$(cat "${WORKDIR}/go-version")
GO_TAR="go${GO_VERSION}.${GOOS}-${GOARCH}.tar.gz"
GO_SRC="https://go.dev/dl/${GO_TAR}"
GO_PATH="${WORKDIR}/go"
DL_PATH="${CACHE_PATH}/downloads/${GO_TAR}"

echo "GOOS: ${GOOS}"
echo "GOARCH: ${GOARCH}"
echo "GO_PATH: ${GO_PATH}"

function usage() {
    echo
    echo "To use the newly installed go binary:"
    echo "  - add ${GO_PATH}/bin to your PATH environment variable."
    echo "  - set ${GO_PATH} as \$GOROOT environment"

    echo
    echo "To use go in the current terminal session, execute the following: "
    echo "      export PATH=\"${GO_PATH}/bin:\$PATH\""
    echo "      export GOROOT=\"${GO_PATH}\""
}

if [[ -d "${GO_PATH}" ]]; then
    echo "Go is already installed" && usage
    exit 0
else
    echo "Downloading ${GO_SRC}..." && \
        mkdir -p "$(dirname $DL_PATH)" && \
        wget --progress=bar:force:noscroll -O "${DL_PATH}" "${GO_SRC}" && \
        echo "Extracting ${GO_TAR} to ${GO_PATH}..." && \
        mkdir -p "${GO_PATH}" && \
        rm -rf ${GO_PATH} && tar -C $(dirname $GO_PATH) -xzf "${DL_PATH}" && \
        echo "Installed Go ${GO_VERSION} to ${GO_PATH}" && usage
fi

cd $WORKDIR
