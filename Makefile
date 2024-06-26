BIN_DIR := bin
BIN_NAME := app

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

.PHONY: all clean build run test crossbuild goenv env

goenv:
	echo $(GOOS) $(GOARCH)

all: clean build

lint:
	@golangci-lint run

lint-fix:
	@golangci-lint run --fix

vet:
	@go vet $(go list ./... | grep -v /vendor/)

build: clean
	@go build -ldflags="-w -s" -o $(BIN_DIR)/$(BIN_NAME)

run: clean build
	@$(BIN_DIR)/$(BIN_NAME)

test:
	@go test -v ./...

clean:
	@rm -rf $(BIN_DIR)

crossbuild: clean $(BIN_DIR)
	@GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(BIN_DIR)/$(BIN_NAME)_linux_amd64
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o $(BIN_DIR)/$(BIN_NAME)_darwin_amd64
	@GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o $(BIN_DIR)/$(BIN_NAME)_windows_amd64.exe

$(BIN_DIR):
	@mkdir -p $(BIN_DIR)%