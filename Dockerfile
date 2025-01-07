FROM ubuntu:24.04 AS downloads

ARG OPENWRT_VERSION="23.05.5"
ARG OPENWRT_TARGET="x86/64"
ARG OPENWRT_ARCH="x86_64"

RUN apt-get update && apt-get install -y \
    wget tar gzip

# Download OpenWrt rootfs
RUN wget \
    "https://downloads.openwrt.org/releases/${OPENWRT_VERSION}/targets/${OPENWRT_TARGET}/openwrt-${OPENWRT_VERSION}-x86-64-rootfs.tar.gz" \
    -O /openwrt.tar.gz && \
    mkdir /rootfs && \
    tar -xf /openwrt.tar.gz -C /rootfs

# Download specific go version
COPY .go-version .go-version
RUN mkdir -p /packages
RUN export GO_VERSION="$(cat .go-version)" && \
        wget \
        "https://github.com/flarehotspot/golang-releases/releases/download/v${GO_VERSION}/golang_${GO_VERSION}_${OPENWRT_ARCH}.ipk" \
        -O "/packages/golang_${GO_VERSION}-${OPENWRT_ARCH}.ipk" --progress=dot:mega && \
        wget \
        "https://github.com/flarehotspot/golang-releases/releases/download/v${GO_VERSION}/golang-src_${GO_VERSION}_${OPENWRT_ARCH}.ipk" \
        -O "/packages/golang-src_${GO_VERSION}-${OPENWRT_ARCH}.ipk" --progress=dot:mega

# --- Start OpenWrt configuration --------------------
FROM scratch

ENV PATH="$PATH:/home/openwrt/go/bin"
ENV GOCACHE=/app/.tmp/go/cache
ENV GOMODCACHE=/app/.tmp/go/mod

COPY --from=downloads /rootfs /
COPY --from=downloads /packages /packages

RUN mkdir -p /var/lock

RUN opkg update && opkg install \
        sudo make shadow-useradd gcc ar shadow-su

RUN opkg install /packages/*.ipk && \
        rm -rf /packages

# Fix gcc ld errors
RUN ar -rc /usr/lib/libpthread.a && \
        ar -rc /usr/lib/libresolv.a && \
        ar -rc /usr/lib/libdl.a

# Run and own only the runtime files as a non-root user for security
RUN useradd openwrt --create-home --shell /bin/sh
USER openwrt
WORKDIR /app

# Install additional tools
COPY ./scripts/install-tools.sh .
RUN ./install-tools.sh

USER root

ENTRYPOINT ["./core/build/devkit/extras/scripts/entrypoint.sh"]

# Watch and recompile server on file change
CMD export PATH=$PATH:/home/openwrt/go/bin; \
    cp go.work.default go.work && \
    go run --tags=dev ./core/cmd/sync-versions/main.go && \
    reflex \
        -r '\.(go|templ|js|css|json)$' \
        -R 'assets\/dist\/.*' \
        -R 'db/sqlc/.*' \
        -R '^config\/.*\.json$' \
        -R 'node_modules' \
        -R '_templ\.go$' \
        -R '(.*)mono\.go' \
        -R '\.tmp\/*.' \
        -R '^output\/*.' \
        -R '^bin\/*.' \
        -R 'core\/main\.go' \
        -R 'plugins\/system\/.*\/main\.go$' \
        -R 'plugins\/local\/.*\/main\.go$' \
        -R 'plugins\/installed\/.*' \
        -R 'plugins\/update\/.*' \
        -R 'plugins\/backup\/.*' \
        -s -- sh -c './start-dev.sh' -v
