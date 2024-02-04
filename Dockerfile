
FROM ubuntu:22.04

ENV BUILD_TAGS=dev

WORKDIR /build

RUN apt-get update && apt-get install -y \
        wget golang-go nodejs npm ca-certificates openssl

COPY . .

RUN npm install -g n && n 20 && \
        hash -r && \
        npm install && \
        node ./build/install-go.js && \
        rm -rf plugins && \
        node ./build/make-go.work.js

ENV PATH=/build/go/bin:${PATH}
ENV GOROOT=/build/go
ENV GOPATH=/build

RUN echo "Using go: $(which go)" && \
        echo "Using go version: $(go version)" && \
        go build --buildmode=plugin -ldflags="-s -w" --tags="${BUILD_TAGS}" -trimpath -o /root/plugin.so core/main.go

FROM scratch
COPY --from=0 /root/plugin.so /root/plugin.so

CMD ["echo", "hello"]

