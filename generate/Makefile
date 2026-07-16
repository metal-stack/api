.PHONY: generate
generate:
	# we remove files explicitly to ensure files really get generated
	rm -f ../go/permissions/servicepermissions.go
	rm -f ../go/client/client.go

	go run ./generate.go