default:
	go run ./core/devkit/cli/flare.go fix-workspace
	go run ./core/internal/cli/flare-internal.go make-mono
	go run ./core/internal/main/main_mono.go


docs-build:
	cd core/sdk/mkdocs && mkdocs build

docs-serve:
	cd core/sdk/mkdocs && mkdocs serve
