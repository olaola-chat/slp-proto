
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOFMT=gofmt
SHELL=/bin/bash #解决ubuntu下编译报错问题

fmt:
	$(GOFMT) -l -s -w .

lint:
	golangci-lint run ./...
	go vet ./...

test:
	$(GOTEST) -v -count=1 ./... 


.PHONY: dao
dao:
	./gen_db.sh xianshi xs_user_profile

.PHONY: cli
cli:
	protoc --go_out=paths=source_relative:./protoc-gen-rbp-rpc/proto protoc-gen-rbp-rpc/proto/option.proto
	go install ./protoc-gen-rbp-rpc


.PHONY: rpc
rpc: cli
	protoc --go_out=paths=source_relative:./gen_pb/rpc/user -I=./protoc-gen-rbp-rpc/proto --proto_path=proto/User user_profile_message.proto
	protoc --rbp-rpc_out=paths=source_relative:./rpcclient/User -I=./protoc-gen-rbp-rpc/proto --proto_path=proto/User user_profile_service.proto

