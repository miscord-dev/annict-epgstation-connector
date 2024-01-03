.PHONY: go-generate
go-generate:
	go generate ./...

.PHONY: generate
generate: go-generate

