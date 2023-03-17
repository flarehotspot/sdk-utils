default:
	rm -rf .cache
	go run -tags=dev -race main/main_dev.go

arm:
	rm -rf .cache
	go run -tags=dev main/main_dev.go

prod:
	rm -rf .cache public
	cd core && make prod
	cd main && make prod
	cd ./plugins/default-theme && make prod
	cd ./plugins/wifi-hotspot && make prod
	cd ./plugins/wired-coinslot && make prod
	cd main && make prod
	./main/app
