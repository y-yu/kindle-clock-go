ALL_PROTO_SRC = $(wildcard ./infra/cache/proto/*.proto)

.PHONY: gen_proto
gen_proto: $(ALL_PROTO_SRC)
	protoc --go_out=. --go_opt=paths=source_relative $^

.PHONY: wire_gen
wire_gen: inject/wire.go
	wire gen $<

GIT_COMMIT_HASH ?= $(shell git rev-parse HEAD)

.PHONY: air
air:
	GIT_COMMIT_HASH=$(GIT_COMMIT_HASH) go tool air