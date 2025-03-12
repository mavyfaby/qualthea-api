BINARY_NAME=qualthea-api

build:
	go mod tidy
	go build -ldflags "-s -w" -o bin/$(BINARY_NAME) cmd/app/main.go

clean:
	go clean
	rm -rf bin
