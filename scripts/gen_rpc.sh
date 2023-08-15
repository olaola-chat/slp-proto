#!/bin/bash


cd `dirname $0`
cd ..

export PROTOC="protoc --go_out=${GOPATH}/src -I=./protoc-gen-rbp-rpc/proto -I=./proto -I=./gen_proto"
export PROTO_SERV="protoc --rbp-rpc_out=${GOPATH}/src -I=./protoc-gen-rbp-rpc/proto -I=./proto -I=./gen_proto"

${PROTOC} rpc/user/user_profile_message.proto
${PROTO_SERV} rpc/user/user_profile_service.proto

${PROTOC} rpc/consume/consume_money_message.proto
${PROTO_SERV} rpc/consume/consume_money_service.proto

${PROTOC} rpc/room/room_info_message.proto
${PROTO_SERV} rpc/room/room_info_service.proto



