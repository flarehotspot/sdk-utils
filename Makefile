PLUGINS = com.adopisoft.basic-flare-theme com.flarego.basic-net-mgr com.flarego.basic-system-account

openwrt: export CGO_ENABLED=1
plugin: export CGO_ENABLED=1

default: clean devmono

devmono: clean
	cd main && make
	cp -r ./plugins/* ./vendor
	./link-resources.sh
	./main/app

plugin: clean
	cd core && make plugin
	cd main && make plugin
	./plugins-action.sh "make plugin"
	cp -r ./plugins/* ./vendor
	./link-resources.sh
	./main/app

openwrt: clean
	ar -rc /usr/lib/libpthread.a
	ar -rc /usr/lib/libresolv.a
	ar -rc /usr/lib/libdl.a
	cd core && make plugin_prod
	cd main && make plugin
	ash ./plugins-action.sh "make plugin"
	cp -r ./plugins/* ./vendor
	ash ./link-resources.sh
	./main/app

sync:
	scp -O -r $(PWD)/core root@$(remote):/root/flarehotspot

sync_all:
	scp -O -r $(PWD) root@$(remote):/root/flarehotspot

clean:
	rm -rf .tmp .cache/views public *.app
	rm -rf ./vendor && mkdir ./vendor
	find . -name "*.so" -type f -delete
	find . -name "*.app" -type f -delete

pull:
	cd core && git pull &
	cd hardware-db && git pull &
	./plugins-action.sh "git pull" &
	git pull &

push:
	cd core && git push &
	cd hardware-db && git push &
	./plugins-action.sh "git push" &
	git push &

checkout_main:
	cd core && git checkout main &
	cd hardware-db && git checkout main
	./plugins-action.sh "git checkout main" &

profile:
	wrk -d20s http://localhost:3000 &
	go tool pprof --seconds 15 -web http://localhost:3000/debug/pprof/profile &

heap:
	wrk -d20s http://localhost:3000 &
	go tool pprof --seconds 15 -web http://localhost:3000/debug/pprof/heap &
