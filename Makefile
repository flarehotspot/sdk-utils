build: export CGO_ENABLED=0
prod: export CGO_ENABLED=0

newifi_d2: export GOOS=linux
newifi_d2: export GOARCH=mipsle

# build_mips: export CGO_ENABLED=0
build_mips: export GOOS=linux
build_mips: export GOARCH=mips
build_mips: export GCCGO=/usr/bin/mips-linux-gnu-gccgo
build_mips: export LD_LIBRARY_PATH=/opt/gccgo/lib64

openwrt_x86: export CGO_ENABLED=0

default: clean
	go build -race -ldflags="-s -w" -o flarehotspot.app -tags="mono dev" main/main_mono.go
	./flarehotspot.app

build: clean
	go build -ldflags="-s -w" -trimpath -o flarehotspot.app -tags="mono dev" main/main_mono.go


build_mips: clean
	/opt/gccgo/bin/go build -tags="mono dev" -compiler=gccgo -gccgoflags -Wl,-R,/opt/gccgo/lib64 -o flarehotspot.app main/main_mono.go

newifi_d2:
	go build -ldflags="-s -w" -trimpath -o flarehotspot.app -tags="mono dev" main/main_mono.go
	./flarehotspot.app

x86:
	go build -ldflags="-s -w" -trimpath -o flarehotspot.app -tags="mono dev" main/main_mono.go
	./flarehotspot.app

sync:
	scp -O -r $(PWD)/core root@$(remote):/root/flarehotspot
	scp -O -r $(PWD)/goutils root@$(remote):/root/flarehotspot
	scp -O -r $(PWD)/sdk root@$(remote):/root/flarehotspot

sync_all:
	scp -O -r $(PWD) root@$(remote):/root/flarehotspot

plugin:
	rm -rf .cache public
	cd core && make plugin
	cd ./plugins/default-theme && make plugin
	cd ./plugins/wifi-hotspot && make plugin
	cd ./plugins/wired-coinslot && make plugin
	cd main && make plugin
	./main/app

clean:
	rm -rf .cache public *.app
	find . -name "*.so" -type f -delete

pull:
	cd core && git pull &
	cd sdk && git pull &
	cd goutils && git pull &
	cd plugins/default-theme && git pull &
	cd plugins/wifi-hotspot && git pull &
	cd plugins/wired-coinslot && git pull &
	git pull

checkout_main:
	cd core && git checkout main &
	cd goutils && git checkout main &
	cd sdk && git checkout main &
	cd plugins/default-theme && git checkout main &
	cd plugins/wifi-hotspot && git checkout main &
	cd plugins/wired-coinslot && git checkout main &
