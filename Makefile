
SHA := $(shell git rev-parse --short=8 HEAD)
GITVERSION := $(shell git describe --long --all)
BUILDDATE := $(shell date -Iseconds)
VERSION := $(or ${VERSION},$(shell git describe --tags --exact-match 2> /dev/null || git symbolic-ref -q --short HEAD || git rev-parse --short HEAD))
LOCALBIN ?= $(shell pwd)/bin
PROTOC_GEN_CONNECPY ?= $(LOCALBIN)/protoc-gen-connecpy
PROTOC_GEN_CONNECPY_VERSION ?= latest

all: proto generate test

release: proto generate test

.PHONY: proto
proto: protolint protoc-gen-connecpy
	$(MAKE) -C go clean
	$(MAKE) -C python clean
	$(MAKE) -C proto protoc
	$(MAKE) -C python generate

.PHONY: protolint
protolint:
	$(MAKE) -C proto protolint

.PHONY: generate
generate:
	$(MAKE) -C generate generate

.PHONY: test
test:
	$(MAKE) -C go test

.PHONY: protoc-gen-connecpy
protoc-gen-connecpy: $(PROTOC_GEN_CONNECPY)
$(PROTOC_GEN_CONNECPY): $(LOCALBIN)
	mkdir -p bin
	GOBIN=$(LOCALBIN) go install github.com/i2y/connecpy/protoc-gen-connecpy@$(PROTOC_GEN_CONNECPY_VERSION)
