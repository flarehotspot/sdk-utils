default: devmono

devplugin:
	cd core && make plugin
	node ./run.js

devmono:
	npm run serve:mono

