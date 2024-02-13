# Description: Dockerfile for building the development kit (devkit).

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
        wget curl golang-go ca-certificates openssl

COPY ./build/setup_nodejs_16.x .

RUN bash ./setup_nodejs_16.x && \
        apt-get install -y nodejs

ARG NODE_ENV=production
ARG DEVKIT_BUILD=1

ENV NODE_ENV=${NODE_ENV}
ENV DEVKIT_BUILD=${DEVKIT_BUILD}
ENV GO_CUSTOM_PATH=/build/go
ENV PATH=${GO_CUSTOM_PATH}/bin:${PATH}

WORKDIR /build

COPY package.json .
COPY package-lock.json .
RUN npm install

COPY . .

RUN rm -rf ./plugins && \
        node ./build/install-go.js && \
        node ./build/make-go.work.js && \
        node ./build/build-core.js

FROM scratch
COPY --from=0 /build/core/plugin.so /plugin.so

CMD ["echo", "hello"]

