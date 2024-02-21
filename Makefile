default:
	go run ./core/devkit/cli/flare.go install-go ./go
	docker compose up --build


build-docs:
	cd core/sdk/mkdocs && mkdocs build

serve-docs:
	cd core/sdk/mkdocs && mkdocs serve
