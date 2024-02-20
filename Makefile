default:
	go run ./core/devkit/cli/flare.go install-go ./go
	docker compose up --build


build-docs:
	cd core/sdk/docs && zola build

serve-docs:
	cd core/sdk/docs && zola serve
