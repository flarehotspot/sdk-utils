default: devmono

devplugin:
	cd core && make plugin
	./run.sh

devmono:
	node ./build-mono.js && go run -tags="dev mono" ./main/main_mono.go

