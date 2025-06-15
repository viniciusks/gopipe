.PHONY: build test lint

# Variáveis de ambiente
VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse --short HEAD)
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

build:
	go build -ldflags "-X 'main.tag=$(VERSION)' -X 'main.commit=$(COMMIT)' -X 'main.date=$(DATE)'" -o gopipe main.go

test:
	go test ./...

lint:
	golangci-lint run

clean:
	rm -rf gopipe