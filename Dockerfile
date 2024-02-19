FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
        wget curl ca-certificates gcc golang-go git

ENV GO_CUSTOM_PATH=/build/go
ENV GO_ENV=development
ENV GO_TAGS="dev mono"
ENV PATH=${GO_CUSTOM_PATH}/bin:${PATH}
ENV FLARE="./core/devkit/cli/flare.go"
ENV FLARE_INT="./core/internal/cli/flare-internal.go"

WORKDIR /build

COPY ./core/go-version  ./core/go-version
COPY ./core/go.mod      ./core/go.mod
COPY ./core/devkit/     ./core/devkit/
COPY ./core/sdk/        ./core/sdk/
COPY ./main/go.mod      ./main/go.mod
COPY ./go.work.default  ./go.work

RUN go run $FLARE install-go

CMD go run $FLARE_INT make-mono && \
    go run $FLARE_INT server

# RUN go build -o ./bin/flare ./core/devkit/cli/flare.go
# RUN go build -o ./bin/flare-internal ./core/internal/cli/flare-internal.go
# RUN ./bin/flare install-go

# CMD ./bin/flare fix-workspace && \
#     ./bin/flare-internal make-mono && \
#     ./bin/flare-internal server

