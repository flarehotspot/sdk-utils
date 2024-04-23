default:
	go run ./core/devkit/cli/flare.go install-go ./go
	docker compose up --build

server:
	go run ./core/internal/cli/flare-internal.go server

openwrt:
	go build -tags="mono staging" -o ./bin/debug-server ./main/main.go

docs-build:
	cd sdk/mkdocs && mkdocs build

docs-serve:
	cd sdk/mkdocs && mkdocs serve
