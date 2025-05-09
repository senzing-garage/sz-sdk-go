# Makefile for Go project

# Detect the operating system and architecture.

include makefiles/osdetect.mk

# -----------------------------------------------------------------------------
# Variables
# -----------------------------------------------------------------------------

# "Simple expanded" variables (':=')

# PROGRAM_NAME is the name of the GIT repository.
PROGRAM_NAME := $(shell basename `git rev-parse --show-toplevel`)
MAKEFILE_PATH := $(abspath $(firstword $(MAKEFILE_LIST)))
MAKEFILE_DIRECTORY := $(shell dirname $(MAKEFILE_PATH))
TARGET_DIRECTORY := $(MAKEFILE_DIRECTORY)/target
DIST_DIRECTORY := $(MAKEFILE_DIRECTORY)/dist
BUILD_TAG := $(shell git describe --always --tags --abbrev=0  | sed 's/v//')
BUILD_ITERATION := $(shell git log $(BUILD_TAG)..HEAD --oneline | wc -l | sed 's/^ *//')
BUILD_VERSION := $(shell git describe --always --tags --abbrev=0 --dirty  | sed 's/v//')
DOCKER_CONTAINER_NAME := $(PROGRAM_NAME)
DOCKER_IMAGE_NAME := senzing/$(PROGRAM_NAME)
DOCKER_BUILD_IMAGE_NAME := $(DOCKER_IMAGE_NAME)-build
DOCKER_SUT_IMAGE_NAME := $(PROGRAM_NAME)_sut
GIT_REMOTE_URL := $(shell git config --get remote.origin.url)
GIT_REPOSITORY_NAME := $(shell basename `git rev-parse --show-toplevel`)
GIT_VERSION := $(shell git describe --always --tags --long --dirty | sed -e 's/\-0//' -e 's/\-g.......//')
GO_PACKAGE_NAME := $(shell echo $(GIT_REMOTE_URL) | sed -e 's|^git@github.com:|github.com/|' -e 's|\.git$$||' -e 's|Senzing|senzing|')
PATH := $(MAKEFILE_DIRECTORY)/bin:$(PATH)

# Recursive assignment ('=')

GO_OSARCH = $(subst /, ,$@)
GO_OS = $(word 1, $(GO_OSARCH))
GO_ARCH = $(word 2, $(GO_OSARCH))

# Conditional assignment. ('?=')
# Can be overridden with "export"
# Example: "export LD_LIBRARY_PATH=/path/to/my/senzing/er/lib"

DOCKER_IMAGE_TAG ?= $(GIT_REPOSITORY_NAME):$(GIT_VERSION)
GOBIN ?= $(shell go env GOPATH)/bin
LD_LIBRARY_PATH ?= /opt/senzing/er/lib

# Export environment variables.

.EXPORT_ALL_VARIABLES:

# -----------------------------------------------------------------------------
# The first "make" target runs as default.
# -----------------------------------------------------------------------------

.PHONY: default
default: help

# -----------------------------------------------------------------------------
# Operating System / Architecture targets
# -----------------------------------------------------------------------------

-include makefiles/$(OSTYPE).mk
-include makefiles/$(OSTYPE)_$(OSARCH).mk


.PHONY: hello-world
hello-world: hello-world-osarch-specific

# -----------------------------------------------------------------------------
# Dependency management
# -----------------------------------------------------------------------------

.PHONY: dependencies-for-development
dependencies-for-development: dependencies-for-development-osarch-specific
	@go install github.com/daixiang0/gci@latest
	@go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest
	@go install github.com/vladopajic/go-test-coverage/v2@latest
	@go install golang.org/x/tools/cmd/godoc@latest
	@docker-compose pull 2>/dev/null || true
	@sudo npm install -g cspell@latest


.PHONY: dependencies
dependencies:
	@go get -u ./...
	@go get -t -u ./...
	@go mod tidy

# -----------------------------------------------------------------------------
# Generate tests
# -----------------------------------------------------------------------------

.PHONY: generate-tests
generate-tests: generate_senzing_unmarshal_test


.PHONY: generate_senzing_unmarshal_test
generate_senzing_unmarshal_test:
	@rm ./response/response_test.go || true
	@./bin/generate_response_response_test.py


.PHONY: verify
verify:
	@rm ./bin/response-test-cases-verified.json || true
	@./bin/verify_response_test_cases.py

# -----------------------------------------------------------------------------
# Setup
# -----------------------------------------------------------------------------

.PHONY: setup
setup: generate-tests setup-osarch-specific

# -----------------------------------------------------------------------------
# Lint
# -----------------------------------------------------------------------------

.PHONY: lint
lint: golangci-lint govulncheck cspell

# -----------------------------------------------------------------------------
# Build
# -----------------------------------------------------------------------------

PLATFORMS := darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64 windows/arm64
$(PLATFORMS):
	$(info Building $(TARGET_DIRECTORY)/$(GO_OS)-$(GO_ARCH)/$(PROGRAM_NAME))
	@GOOS=$(GO_OS) GOARCH=$(GO_ARCH) go build -o $(TARGET_DIRECTORY)/$(GO_OS)-$(GO_ARCH)/$(PROGRAM_NAME)


.PHONY: build
build: build-osarch-specific

# -----------------------------------------------------------------------------
# Run
# -----------------------------------------------------------------------------

.PHONY: run
run: run-osarch-specific

# -----------------------------------------------------------------------------
# Test
# -----------------------------------------------------------------------------

.PHONY: test
test: test-osarch-specific

# -----------------------------------------------------------------------------
# Coverage
# -----------------------------------------------------------------------------

.PHONY: coverage
coverage: coverage-osarch-specific


.PHONY: check-coverage
check-coverage: export SENZING_LOG_LEVEL=TRACE
check-coverage:
	@go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...
	@${GOBIN}/go-test-coverage --config=.github/coverage/testcoverage.yaml

# -----------------------------------------------------------------------------
# Documentation
# -----------------------------------------------------------------------------

.PHONY: documentation
documentation: documentation-osarch-specific

# -----------------------------------------------------------------------------
# Clean
# -----------------------------------------------------------------------------

.PHONY: clean
clean: clean-osarch-specific
	@go clean -cache
	@go clean -testcache

# -----------------------------------------------------------------------------
# Utility targets
# -----------------------------------------------------------------------------

.PHONY: help
help:
	$(info Build $(PROGRAM_NAME) version $(BUILD_VERSION)-$(BUILD_ITERATION))
	$(info Makefile targets:)
	@$(MAKE) -pRrq -f $(firstword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs


.PHONY: print-make-variables
print-make-variables:
	@$(foreach V,$(sort $(.VARIABLES)), \
		$(if $(filter-out environment% default automatic, \
		$(origin $V)),$(info $V=$($V) ($(value $V)))))


.PHONY: update-pkg-cache
update-pkg-cache:
	@GOPROXY=https://proxy.golang.org GO111MODULE=on \
		go get $(GO_PACKAGE_NAME)@$(BUILD_TAG)

# -----------------------------------------------------------------------------
# Specific programs
# -----------------------------------------------------------------------------

.PHONY: cspell
cspell:
	@cspell .


.PHONY: golangci-lint
golangci-lint:
	@${GOBIN}/golangci-lint run --config=.github/linters/.golangci.yaml


.PHONY: govulncheck
govulncheck:
	@${GOBIN}/govulncheck ./...
