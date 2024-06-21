FROM ubuntu:24.04

RUN apt-get update && \
        apt-get install -y \
        wget curl gcc golang-go git ca-certificates

ENV GOPATH=/opt/go
ENV GOCACHE=/build/.tmp/gocache
ENV GO_CUSTOM_PATH=/build/go
ENV PATH=${GO_CUSTOM_PATH}/bin:${PATH}
ENV PATH=${PATH}:/opt/go/bin

WORKDIR /build

RUN go install github.com/cespare/reflex@latest

CMD cp go.work.default go.work && \
    go run --tags=dev ./core/internal/cli/main.go install-go && \
    go run --tags=dev ./core/cmd/make-mono/main.go && \
    reflex \
        -r '\.go$' \
        -R 'core\/main\.go' \
        -R 'plugins\/.*\/main\.go' \
        -R '(.*)mono\.go' \
        -s -- sh -c './start.sh' -v
