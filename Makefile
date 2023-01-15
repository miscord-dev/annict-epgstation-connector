epgstation/api.gen.go: epgstation/schema.json
	oapi-codegen -package epgstation schema/schema.json > epgstation/api.gen.go

.PHONY: go-generate
go-generate:
	go generate ./...

.PHONY: generate
generate: epgstation/api.gen.go go-generate

