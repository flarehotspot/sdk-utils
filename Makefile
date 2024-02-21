default:
	go run ./core/devkit/cli/flare.go install-go ./go
	docker compose up --build


docs-build:
	cd core/sdk/mkdocs && mkdocs build

docs-serve:
	cd core/sdk/mkdocs && mkdocs serve
