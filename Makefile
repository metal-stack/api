
SHA := $(shell git rev-parse --short=8 HEAD)
GITVERSION := $(shell git describe --long --all)
BUILDDATE := $(shell date -Iseconds)
VERSION := $(or ${VERSION},$(shell git describe --tags --exact-match 2> /dev/null || git symbolic-ref -q --short HEAD || git rev-parse --short HEAD))
LOCALBIN ?= $(shell pwd)/bin
PROTOC_GEN_CONNECPY ?= $(LOCALBIN)/protoc-gen-connecpy
PROTOC_GEN_CONNECPY_VERSION ?= 2.3.0

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
protoc-gen-connecpy:
	mkdir -p $(LOCALBIN)
	rm -f $(LOCALBIN)/protoc-gen-connecpy
	curl -s -L https://github.com/i2y/connecpy/releases/download/v$(PROTOC_GEN_CONNECPY_VERSION)/protoc-gen-connecpy_$(PROTOC_GEN_CONNECPY_VERSION)_linux_amd64.tar.gz | tar -xzf - protoc-gen-connecpy
	mv protoc-gen-connecpy $(LOCALBIN)/protoc-gen-connecpy
