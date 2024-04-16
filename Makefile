default:
	go run ./core/devkit/cli/flare.go install-go ./go
	docker compose up --build

docs-build:
	cd sdk/mkdocs && mkdocs build

docs-serve:
	cd sdk/mkdocs && mkdocs serve
