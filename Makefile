default:
	go run ./core/devkit/cli/flare.go install-go ./go
	docker compose up --build

server:
	go run ./core/internal/cli/flare-internal.go server

docs-build:
	cd sdk/mkdocs && mkdocs build

docs-serve:
	cd sdk/mkdocs && mkdocs serve
