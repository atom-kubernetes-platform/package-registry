.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golangci-lint run
.PHONY:lint

vet: lint
	go vet ./...
.PHONY:vet

test: vet
	go test -v ./...
.PHONY:test

build: test
	go build .
.PHONY:build
