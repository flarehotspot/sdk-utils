default: export CGO_ENABLED=0
build: export CGO_ENABLED=0
prod: export CGO_ENABLED=0

default: clean
	go build -ldflags="-s -w" -o flarehotspot.app -tags="mono dev" main/main_mono.go
	./flarehotspot.app

build: clean
	go build -ldflags="-s -w" -trimpath -o flarehotspot.app -tags="mono dev" main/main_mono.go

prod: clean
	go build -ldflags="-s -w" -trimpath -o flarehotspot.app -tags="mono" main/main_mono.go

serve_prod: prod
	./app

plugin:
	rm -rf .cache public
	cd core && make plugin
	cd ./plugins/default-theme && make plugin
	cd ./plugins/wifi-hotspot && make plugin
	cd ./plugins/wired-coinslot && make plugin
	cd main && make plugin
	./main/app

clean:
	rm -rf .cache public
