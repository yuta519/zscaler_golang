.PHONY: install test build-zia

install:
	go mod tidy

test:
	go test

build:
	go mod tidy
	go build -o zia cmd/zia/main.go
