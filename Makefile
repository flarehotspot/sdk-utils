default:
	go run ./core/devkit/cli/flare.go fix-workspace
	go run ./core/internal/cli/flare-internal.go make-mono
	go run ./core/internal/main/main_mono.go


build-docs:
	cd core/sdk/docs && zola build

serve-docs:
	cd core/sdk/docs && zola serve
