default:
	go run ./core/devkit/cli/flare.go install-go ./go
	docker compose up --build

server-dev:
	./run-dev.sh

openwrt:
	go build -tags="mono staging" -o ./bin/debug-server ./main/main.go

docs-build:
	cd sdk/mkdocs && mkdocs build

docs-serve:
	cd sdk/mkdocs && mkdocs serve
