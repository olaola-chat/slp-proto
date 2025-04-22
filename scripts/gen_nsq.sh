#!/bin/bash


cd `dirname $0`
cd ..

export PROTOC="protoc --go_out=${GOPATH}/src -I=./protoc-gen-slp-rpc/proto -I=./proto -I=./gen_proto"
export PROTO_SERV="protoc --slp-rpc_out=${GOPATH}/src -I=./protoc-gen-slp-rpc/proto -I=./proto -I=./gen_proto"

${PROTOC} nsq/consume/common_consume.proto
${PROTOC} nsq/room/room_consume_package.proto



