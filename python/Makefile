.PHONY: clean
clean:
	rm -rf metalstack/admin
	rm -rf metalstack/api
	rm -rf metalstack/infra

.PHONY: generate
generate:
	@find metalstack -type d -exec touch {}/__init__.py \;
	@find buf -type d -exec touch {}/__init__.py \;
