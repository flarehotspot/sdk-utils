FROM ubuntu:22.04

RUN apt-get update && \
        apt-get install -y \
        wget curl gcc golang-go git ca-certificates

ENV GOCACHE=/build/.tmp/gocache
ENV GO_CUSTOM_PATH=/build/go
ENV GO_ENV=development
ENV PATH=${GO_CUSTOM_PATH}/bin:${PATH}
ENV FLARE="./core/devkit/cli/flare.go"
ENV FLARE_INT="./core/internal/cli/flare-internal.go"

WORKDIR /build

RUN go install github.com/cespare/reflex@latest

CMD cp go.work.default go.work && \
    go run $FLARE install-go && \
    go run $FLARE_INT make-mono && \
    $HOME/go/bin/reflex \
        -r '\.go$' \
        -R 'core\/main\.go' \
        -R 'plugins\/.*\/main\.go' \
        -R 'mono\.go' \
        -s -- sh -c './start.sh' -v
