
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOFMT=gofmt
SHELL=/bin/bash #解决ubuntu下编译报错问题
PWD=$(shell pwd)
PROTOC=protoc --go_out=${GOPATH}/src -I=./protoc-gen-slp-rpc/proto -I=./proto -I=./gen_proto
PROTO_SERV=protoc --slp-rpc_out=${GOPATH}/src -I=./protoc-gen-slp-rpc/proto -I=./proto -I=./gen_proto

fmt:
	$(GOFMT) -l -s -w .

lint:
	golangci-lint run ./...
	go vet ./...

test:
	GF_GCFG_PATH=$(PWD) $(GOTEST) -v -count=1 ./... 

.PHONY: dao
dao:
	./scripts/gen_dao_xianshi.sh
	./scripts/gen_dao_config.sh
	./scripts/gen_dao_activity.sh
	./scripts/gen_dao_functor.sh

.PHONY: cli
cli:
	${PROTOC} slp/plugin/option.proto
	go install ./protoc-gen-slp-rpc


.PHONY: rpc
rpc: cli
	./scripts/gen_rpc.sh


.PHONY: nsq
nsq:
	./scripts/gen_nsq.sh


.PHONY: kafka
kafka:
	./scripts/gen_kafka.sh

.PHONY: es
es:
	./scripts/gen_es.sh
