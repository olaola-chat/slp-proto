
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOFMT=gofmt
SHELL=/bin/bash #解决ubuntu下编译报错问题
PWD=$(shell pwd)
PROTOC=protoc --go_out=${GOPATH}/src -I=./protoc-gen-rbp-rpc/proto -I=./proto -I=./gen_proto
PROTO_SERV=protoc --rbp-rpc_out=${GOPATH}/src -I=./protoc-gen-rbp-rpc/proto -I=./proto -I=./gen_proto

fmt:
	$(GOFMT) -l -s -w .

lint:
	golangci-lint run ./...
	go vet ./...

test:
	GF_GCFG_PATH=$(PWD) $(GOTEST) -v -count=1 ./... 

.PHONY: dao
dao:
	./gen_dao_xianshi.sh

.PHONY: cli
cli:
	${PROTOC} rbp/plugin/option.proto
	go install ./protoc-gen-rbp-rpc


.PHONY: rpc
rpc: cli
	${PROTOC} user/user_profile_message.proto
	${PROTO_SERV} user/user_profile_service.proto
	${PROTOC} voice_lover/voice_lover_common.proto
	${PROTOC} voice_lover/voice_lover_message.proto
	${PROTO_SERV} voice_lover/voice_lover_main_service.proto
	${PROTO_SERV} voice_lover/voice_lover_admin_service.proto

