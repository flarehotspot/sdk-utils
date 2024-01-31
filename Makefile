default: devmono

devplugin:
	cd core && make plugin
	./run.sh

devmono:
	npm run serve:mono

