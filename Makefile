SERVICE_NAME=gotorch

.PHONY: fmt
fmt:
	gofmt -w -s .
	goimports -w .
	go clean ./...


GREEN=\033[0;32m # Green
NC=\033[0m # No Color

.PHONY: lint
lint:
	golangci-lint run --config golangci.yml --timeout=10m
	@echo "${GREEN}[✔] golangci-lint OK${NC}"


.PHONY: build
build:
	go build \
		-race \
		-o ./bin/$(SERVICE_NAME) \
		-v ./cmd/main.go

.PHONY: run
run: build
	bin/$(SERVICE_NAME)

