.PHONY: clean
clean:
	rm -rf metalstack

.PHONY: generate
generate:
	@find metalstack -type d -exec touch {}/__init__.py \;
	@find buf -type d -exec touch {}/__init__.py \;
