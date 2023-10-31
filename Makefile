default: devplugin

devplugin:
	cd core && make plugin
	./run.sh

devmono:
	node ./build-mono.js && ./main/main.app

