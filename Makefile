BINARY_NAME=qualthea-api

build:
	go mod tidy
	go build -o bin/$(BINARY_NAME) cmd/app/main.go

clean:
	go clean
	rm -rf bin
