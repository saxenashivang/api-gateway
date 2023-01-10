GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	@protoc -I . --proto_path=../../protos --micro_out=. --go_out=:.  proto/userservice.proto