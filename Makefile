build: export CGO_ENABLED=0
prod: export CGO_ENABLED=0

newifi_d2: export GOOS=linux
newifi_d2: export GOARCH=mipsle

default: clean
	go build -race -ldflags="-s -w" -o flarehotspot.app -tags="mono dev" main/main_mono.go
	./flarehotspot.app

build: clean
	go build -ldflags="-s -w" -trimpath -o flarehotspot.app -tags="mono dev" main/main_mono.go

serve_prod: prod
	./app

newifi_d2:
	go build -ldflags="-s -w" -trimpath -o flarehotspot.app -tags="mono dev" main/main_mono.go
	./flarehotspot.app

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

pull:
	cd core && git pull &
	cd sdk && git pull &
	cd goutils && git pull &
	cd plugins/default-theme && git pull &
	cd plugins/wifi-hotspot && git pull &
	cd plugins/wired-coinslot && git pull &
	git pull
