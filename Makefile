default:
	rm -rf .cache
	go run -tags=dev -race main/main_dev.go
