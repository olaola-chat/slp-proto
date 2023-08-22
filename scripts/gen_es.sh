#!/bin/bash


cd `dirname $0`
cd ..

export PROTOC="protoc --go_out=${GOPATH}/src -I=./protoc-gen-rbp-rpc/proto -I=./proto -I=./gen_proto"
export PROTO_SERV="protoc --rbp-rpc_out=${GOPATH}/src -I=./protoc-gen-rbp-rpc/proto -I=./proto -I=./gen_proto"

${PROTOC} es/room/room_search.proto


