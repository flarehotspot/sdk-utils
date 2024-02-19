# Description: Dockerfile for running the devkit.

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
        wget curl gcc ca-certificates

ENV GO_ENV=development
ENV GO_CUSTOM_PATH=/build/go
ENV PATH=${GO_CUSTOM_PATH}/bin:${PATH}

WORKDIR /build
COPY ./bin ./bin
COPY ./core/go-version ./core/go-version
RUN ./bin/flare install-go

CMD ./bin/flare server

