
FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
        wget curl golang-go ca-certificates openssl

RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash - &&\
        apt-get install -y nodejs

ENV BUILD_TAGS=dev
ENV GO_CUSTOM_PATH=/build/go
ENV PATH=${GO_CUSTOM_PATH}/bin:${PATH}
WORKDIR /build



COPY . .

RUN npm install && \
        node ./build/install-go.js && \
        rm -rf plugins && \
        node ./build/make-go.work.js


RUN echo "Using go: $(which go)" && \
        echo "Using go version: $(go version)" && \
        go build --buildmode=plugin -ldflags="-s -w" --tags="${BUILD_TAGS}" -trimpath -o /root/plugin.so core/main.go

FROM scratch
COPY --from=0 /root/plugin.so /root/plugin.so

CMD ["echo", "hello"]

