default: devplugin

devplugin:
	cd core && make plugin
	./run.sh

devmono:
	cd main && go build --tags="dev mono" -o main.app main_mono.go
	./main/main.app

