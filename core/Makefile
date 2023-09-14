plugin: export CGO_ENABLED=1

plugin:
	go build -buildmode=plugin -tags="dev" -trimpath -o core.so ./main.go

devkit:
	go build -buildmode=plugin -tags="dev" -trimpath -o core.so ./main.go

plugin_prod:
	go build -buildmode=plugin -ldflags="-s -w" -trimpath -o core.so ./main.go

test:
	go test -v ./...

# command for running a single test:
# t=Test_name_of_test_function make testonly
testonly:
	go test -v -run $(t) ./...
