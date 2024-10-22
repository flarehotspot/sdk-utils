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

CMD cp go.work.default go.work && \
    go run --tags=dev ./core/internal/cli/main.go install-go && \
    go run --tags=dev ./core/cmd/sync-versions/main.go && \
    cd tools && ./install.sh && cd .. && \
    reflex \
        -r '\.go$' \
        -R '\.tmp\/*.' \
        -R 'core\/main\.go' \
        -R 'plugins\/installed' \
        -R 'plugins\/system\/.*\/main\.go' \
        -R 'plugins\/local\/.*\/main\.go' \
        -R 'plugins\/update\/.*\/main\.go' \
        -R 'plugins\/backup\/.*\/main\.go' \
        -R '(.*)mono\.go' \
        -R 'tools' \
        -s -- sh -c './start.sh' -v
