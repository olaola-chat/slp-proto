
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOFMT=gofmt
SHELL=/bin/bash #解决ubuntu下编译报错问题
PROTOC=protoc --go_out=${GOPATH}/src -I=./protoc-gen-rbp-rpc/proto -I=./proto -I=./gen_proto
PROTO_SERV=protoc --rbp-rpc_out=${GOPATH}/src -I=./protoc-gen-rbp-rpc/proto -I=./proto -I=./gen_proto

fmt:
	$(GOFMT) -l -s -w .

lint:
	golangci-lint run ./...
	go vet ./...

test:
	$(GOTEST) -v -count=1 ./... 

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

