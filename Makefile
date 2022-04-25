.PHONY: install test build-zia

install:
	go mod tidy

test:
	go test

build:
	go build -o zia cmd/zia/main.go
