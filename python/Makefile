.PHONY: clean
clean:
	rm -rf buf
	rm -rf metalstack/admin
	rm -rf metalstack/api
	rm -rf metalstack/infra
	rm -rf metalstack/cli

.PHONY: generate
generate:
	@find metalstack -type d -exec touch {}/__init__.py \;
	@find buf -type d -exec touch {}/__init__.py \;
