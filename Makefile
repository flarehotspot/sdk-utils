default:
	rm -rf .cache public
	mkdir .cache
	ln -s ${PWD}/default-theme ${PWD}/.cache/theme
	go run -tags=dev -race main/main_dev.go
