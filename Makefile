.PHONY: default
default: help

# Environment variables
NOW = $(shell date +%Y-%m-%dT%H:%M:%S%z)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
COMMIT_DATE = $(shell git show -s --format=%cI)
COMMIT_HASH = $(shell git rev-parse HEAD)
GOLANGCI_LINT_VERSION=v1.42
GOLANGCI_LINT_COMMAND=golangci-lint run --config=.golangci.yml --allow-parallel-runners --timeout 30m -v
PRETTIER_NODE_DOCKER_IMAGE_TAG=16.13.1-alpine3.15
PRETTIER_VERSION=2.5.1

#> help: Show help menu
.PHONY: help
help: Makefile
	@echo
	@echo "Available targets:"
	@echo
	@sed -n 's/^#>//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

#> test: Build and run unit tests
.PHONY: test
test:
	@echo "Testing..."
	@env go test -race $(shell go list ./... | grep -v /vendor/) -count=1

#> go-lint: Lint golang files
.PHONY: go-lint
go-lint:
	golangci-lint run --timeout 30m

#> prettier-check: Check whether files are formatted correctly with Prettier.
.PHONY: prettier-check
prettier-check:
	#> Use Docker to pin Prettier to the same version as the one run via Drone.
	docker run --rm -it \
		-w /app \
		-v ${PWD}:/app:ro \
		node:${PRETTIER_NODE_DOCKER_IMAGE_TAG} \
		/bin/sh -c \
			"npm install --global --save-dev --save-exact prettier@${PRETTIER_VERSION} \
			&& prettier -c ."

#> prettier-write: Format files with Prettier.
.PHONY: prettier-write
prettier-write:
	#> Use Docker to pin Prettier to the same version as the one run via Drone.
	docker run --rm -it \
		-w /app \
		-v ${PWD}:/app \
		node:${PRETTIER_NODE_DOCKER_IMAGE_TAG} \
		/bin/sh -c \
			"npm install --global --save-dev --save-exact prettier@${PRETTIER_VERSION} \
			&& prettier -w ."

# These are all paths that should be built every time they are called as targets directly or indirectly.
.PHONY: \
	lint \
	prettier-check \
	prettier-write \
	test \