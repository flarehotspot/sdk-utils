# Description: Dockerfile for running the devkit.

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

COPY ./build/exec-async.js  ./build/exec-async.js
COPY ./build/install-go.js  ./build/install-go.js
COPY ./core/go-version      ./core/go-version
RUN node ./build/install-go.js

CMD ["npm", "start"]
