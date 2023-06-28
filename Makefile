openwrt: export CGO_ENABLED=1
plugin: export CGO_ENABLED=1

default: clean
	go build -race -ldflags="-s -w" -o flarehotspot.app -tags="mono dev" main/main_mono.go
	./flarehotspot.app

build: clean
	go build -ldflags="-s -w" -trimpath -o flarehotspot.app -tags="mono dev" main/main_mono.go

# openwrt:
	# go build -ldflags="-s -w" -trimpath -o flarehotspot.app -tags="mono prod" main/main_mono.go
	# ./flarehotspot.app

openwrt:
	ar -rc /usr/lib/libpthread.a
	ar -rc /usr/lib/libresolv.a
	ar -rc /usr/lib/libdl.a
	rm -rf .cache public
	cd core && make plugin_prod
	cd ./plugins/flarehotspot-theme && make plugin
	cd ./plugins/wifi-hotspot && make plugin
	cd ./plugins/wired-coinslot && make plugin
	cd ./plugins/basic-system-accounts && make plugin
	cd main && make plugin
	./main/app

sync:
	scp -O -r $(PWD)/core root@$(remote):/root/flarehotspot
	scp -O -r $(PWD)/goutils root@$(remote):/root/flarehotspot
	scp -O -r $(PWD)/sdk root@$(remote):/root/flarehotspot

sync_all:
	scp -O -r $(PWD) root@$(remote):/root/flarehotspot

plugin:
	rm -rf .cache public
	cd core && make plugin
	cd ./plugins/flarehotspot-theme && make plugin
	cd ./plugins/wifi-hotspot && make plugin
	cd ./plugins/wired-coinslot && make plugin
	cd ./plugins/basic-system-accounts && make plugin
	cd main && make plugin
	./main/app

clean:
	rm -rf .cache public *.app
	find . -name "*.so" -type f -delete

pull:
	cd core && git pull &
	cd sdk && git pull &
	cd goutils && git pull &
	cd plugins/flarehotspot-theme && git pull &
	cd plugins/wifi-hotspot && git pull &
	cd plugins/wired-coinslot && git pull &
	cd plugins/basic-system-accounts && git pull &
	git pull

push:
	cd core && git push &
	cd sdk && git push &
	cd goutils && git push &
	cd plugins/flarehotspot-theme && git push &
	cd plugins/wifi-hotspot && git push &
	cd plugins/wired-coinslot && git push &
	cd plugins/basic-system-accounts && git push &
	git push

checkout_main:
	cd core && git checkout main &
	cd goutils && git checkout main &
	cd sdk && git checkout main &
	cd plugins/flarehotspot-theme && git checkout main &
	cd plugins/wifi-hotspot && git checkout main &
	cd plugins/wired-coinslot && git checkout main &
	cd plugins/basic-system-accounts && git checkout main &
