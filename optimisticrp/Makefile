VERSION=$(shell cat VERSION)
SMART_CONTRACTS=$(wildcard store/*.sol)
APP_NAME=mydict

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## Build the container
	solc --abi contracts/* -o contracts/build
	abigen --abi=contracts/build/Optimistic_Roll_In.abi --pkg=contracts --out=contracts/Contracts.go


deploy: ## Build the container
	docker-compose up -d

clean: ## Clean files
	rm -rf ./contracts/build
	rm contracts/*.go
