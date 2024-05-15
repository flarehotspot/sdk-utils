default:
	docker compose up --build

server-dev:
	./run-dev.sh

openwrt:
	go build -tags="mono staging" -o ./bin/debug-server ./main/main.go

docs-build:
	cd sdk/mkdocs && mkdocs build

docs-serve:
	cd sdk/mkdocs && mkdocs serve

devkit:
	go run -tags="dev" ./core/internal/cli/flare-internal.go create-devkit
