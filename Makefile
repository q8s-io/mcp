# If you update this file, please follow
# https://suva.sh/posts/well-documented-makefiles

# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL:=/usr/bin/env bash

.DEFAULT_GOAL := help

# Use GOPROXY environment variable if set
GOPROXY := $(shell go env GOPROXY)
ifeq ($(GOPROXY),)
GOPROXY := https://goproxy.cn
endif
export GOPROXY

# Active module mode, as we use go modules to manage dependencies
export GO111MODULE=on

TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/bin
BIN_DIR := bin

GOLANGCI_LINT := $(TOOLS_BIN_DIR)/golangci-lint

# Define Docker related variables. Releases should modify and double check these vars.
REGISTRY ?= docker.io/q8sio
IMAGE_NAME ?= mcp
CONTROLLER_IMG ?= $(REGISTRY)/$(IMAGE_NAME)
TAG ?= dev
ARCH ?= amd64

## --------------------------------------
## Help
## --------------------------------------

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Testing
## --------------------------------------

.PHONY: test
test: ## Run tests
	debug=false go test ./...

.PHONY: test-integration
test-integration: ## Run integration tests
	CONFIG_FILE=configs/config-dev.yaml go test -v -tags=integration ./test/integration/...

## --------------------------------------
## Tooling Binaries
## --------------------------------------

$(GOLANGCI_LINT): $(TOOLS_DIR)/go.mod # Build golangci-lint from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

## --------------------------------------
## Linting
## --------------------------------------

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Lint codebase
	$(GOLANGCI_LINT) run -v

lint-full: $(GOLANGCI_LINT) ## Run slower linters to detect possible issues
	$(GOLANGCI_LINT) run -v --fast=false

## --------------------------------------
## Generate
## --------------------------------------

.PHONY: modules
modules: ## Runs go mod to ensure proper vendoring.
	go mod tidy
	cd $(TOOLS_DIR); go mod tidy

.PHONY: generate-swagger-json
generate-swagger-json: ## Generate swagger.json file
	go run cmd/mcp-swagger/main.go

## --------------------------------------
## Docker
## --------------------------------------

.PHONY: docker-build
docker-build: ## Build the docker image for controller-manager
	docker build --pull --build-arg ARCH=$(ARCH) . -t $(CONTROLLER_IMG):$(TAG)

.PHONY: docker-push
docker-push: ## Push the docker image
	docker push $(CONTROLLER_IMG):$(TAG)

## --------------------------------------
## Cleanup / Verification
## --------------------------------------

.PHONY: clean
clean: ## Remove all generated files
	$(MAKE) clean-bin

.PHONY: clean-bin
clean-bin: ## Remove all generated binaries
	rm -rf bin
	rm -rf hack/tools/bin

## --------------------------------------
## Run
## --------------------------------------

.PHONY: run-server
run-server: ## Run mcp server
	go run cmd/mcp-server/main.go --config=configs/config-dev.yaml

.PHONY: run-swagger
run-swagger: generate-swagger-json ## Run swagger ui
	docker run -d -p 8081:8080 --rm \
	    -e SWAGGER_JSON=/usr/share/nginx/swagger.json \
	    -v `pwd`/api/openapi-spec/swagger.json:/usr/share/nginx/swagger.json \
	    --name swagger-ui swaggerapi/swagger-ui:3.18.2

.PHONY: run-syncdb
run-syncdb: ## Run sync db
	go run cmd/mcp-syncdb/main.go --config=configs/config-dev.yaml
