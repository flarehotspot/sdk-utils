
FROM ubuntu:22.04

ENV BUILD_TAGS=dev

WORKDIR /build

RUN apt-get update && apt-get install -y \
        wget golang-go ca-certificates openssl

COPY . .

RUN ./install-go.sh
RUN ./go-work.sh

ENV PATH=/build/go/bin:${PATH}
ENV GOROOT=/build/go
ENV GOPATH=/build

RUN echo "Using go: $(which go)" && \
        echo "Using go version: $(go version)" && \
        go build --buildmode=plugin -ldflags="-s -w" --tags="${BUILD_TAGS}" -trimpath -o /root/core.so core/main.go

FROM scratch
COPY --from=0 /root/core.so /root/core.so

CMD ["echo", "hello"]

