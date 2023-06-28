#/bin/sh

ROOT_REPO="github.com/flarehotspot"

export GOPRIVATE="$ROOT_REPO/core"
export GOPRIVATE="$GOPRIVATE,$ROOT_REPO/goutils"
export GOPRIVATE="$GOPRIVATE,$ROOT_REPO/sdk"
export GOPRIVATE="$GOPRIVATE,$ROOT_REPO/default-theme"
export GOPRIVATE="$GOPRIVATE,$ROOT_REPO/wifi-hotspot"
export GOPRIVATE="$GOPRIVATE,$ROOT_REPO/wired-coinslot"
export GOPRIVATE="$GOPRIVATE,$ROOT_REPO/basic-system-account"
export GOPRIVATE="$GOPRIVATE,$ROOT_REPO/basic-net-mgr"

