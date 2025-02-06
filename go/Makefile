.PHONY: clean
clean:
	rm -rf admin/v2/*.pb*.go admin status
	rm -rf api/v2/*.pb*.go admin status
	rm -rf metalstack/api/v2/*.pb*.go admin status
	rm -rf metalstack/admin/v2/*.pb*.go admin status

.PHONY: test
test:
	go test -v ./...
