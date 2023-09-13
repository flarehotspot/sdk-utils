default: clean
	./go-work.sh
	cd main && go build -ldflags="-s -w" -tags="dev" -trimpath -o main.app main.go
	./build-plugins.sh
	cp -r ./plugins/* ./vendor
	./link-resources.sh
	./main/main.app

clean:
	rm -rf .tmp .cache/views public
	rm -rf ./vendor && mkdir ./vendor
	find . -name "plugins/*.so" -type f -delete
	find . -name "*.app" -type f -delete
