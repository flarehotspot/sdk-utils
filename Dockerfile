FROM ubuntu:24.04

RUN apt-get update && apt-get install -y \
    wget tar gzip make gcc git sudo zip

ENV PATH=${PATH}:/root/go/bin
ENV PATH=${PATH}:/usr/local/go/bin

# Install go
COPY .go-version .
RUN wget https://go.dev/dl/go$(cat .go-version).linux-$(dpkg --print-architecture).tar.gz\
        -O golang.tar.gz && \
        rm -rf /usr/local/go && \
        tar -C /usr/local -xzf golang.tar.gz && \
        rm -rf golang.tar.gz

WORKDIR /app

# Install additional tools
COPY ./core/build/devkit/extras/scripts/install-tools.sh .
RUN ./install-tools.sh

# Watch and recompile server on file change
CMD cp go.work.default go.work && \
    go run --tags=dev ./core/cmd/sync-versions/main.go && \
    reflex \
        -r '\.(go|templ|js|css|json)$' \
        -R 'assets\/dist\/.*' \
        -R 'db/sqlc/.*' \
        -R '^config\/.*\.json$' \
        -R 'node_modules' \
        -R '_templ\.go$' \
        -R '\.tmp\/*.' \
        -R '^output\/*.' \
        -R '^bin\/*.' \
        -R 'plugins\/installed\/.*' \
        -R 'plugins\/update\/.*' \
        -R 'plugins\/backup\/.*' \
        -s -- sh -c './start-dev.sh' -v
