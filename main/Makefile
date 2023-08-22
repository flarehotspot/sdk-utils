plugin: export CGO_ENABLED=1

default: build

build:
	go build -ldflags="-s -w" -tags="dev mono" -trimpath -o app main_mono.go

plugin:
	go build -ldflags="-s -w" -tags="dev" -trimpath -o app main.go
