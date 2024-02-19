FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
        wget curl ca-certificates gcc golang-go git

ENV GO_CUSTOM_PATH=/build/go
ENV PATH=${GO_CUSTOM_PATH}/bin:${PATH}

WORKDIR /build
RUN ./bin/flare install-go

COPY . .

RUN ./flare install-go

CMD ["./main/main.app"]
