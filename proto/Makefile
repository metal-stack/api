MAKEFLAGS += --no-print-directory
BUF_VERSION := 1.72.0

.ONESHELL:
_gobuf:
	docker build --tag gobuf - <<-'EOT'
	FROM golang:1.26 AS gobase
	FROM bufbuild/buf:${BUF_VERSION} AS buf

	FROM gobase
	COPY --from=buf /usr/local/bin/buf /usr/local/bin/buf
	EOT

_buf:
	docker run --rm \
		--entrypoint sh \
		-v $(PWD)/..:/workspace \
		-w /workspace/proto \
		gobuf \
		-c "buf $(CMD) && chown -R $(shell id -u):$(shell id -g) /workspace"

.PHONY: protolint
protolint: _gobuf
	@$(MAKE)	_buf	CMD="format -w"
	@$(MAKE)	_buf	CMD="lint -v"
	@$(MAKE)	_buf	CMD="breaking --against 'https://github.com/metal-stack/api.git#branch=main,subdir=proto'"

.PHONY: protoc
protoc: _gobuf
	@$(MAKE)	_buf	CMD="generate -v"
