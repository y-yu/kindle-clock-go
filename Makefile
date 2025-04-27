ALL_PROTO_SRC  = $(wildcard ./infra/cache/proto/*.proto)

.PHONY: gen_proto
gen_proto: $(ALL_PROTO_SRC)
	protoc --go_out=. --go_opt=paths=source_relative $^