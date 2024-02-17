# Description: Dockerfile for running the devkit.

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
        wget curl golang-go ca-certificates openssl

ARG NODE_ENV=production
ARG DEVKIT_BUILD=1

ENV NODE_ENV=${NODE_ENV}
ENV DEVKIT_BUILD=${DEVKIT_BUILD}
ENV GO_CUSTOM_PATH=/build/go
ENV PATH=${GO_CUSTOM_PATH}/bin:${PATH}

WORKDIR /build

COPY . .

RUN ./flare install-go

CMD ["./main/main.app"]
