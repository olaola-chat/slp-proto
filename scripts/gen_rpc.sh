#!/bin/bash


cd `dirname $0`
cd ..

export PROTOC="protoc --go_out=${GOPATH}/src -I=./protoc-gen-rbp-rpc/proto -I=./proto -I=./gen_proto"
export PROTO_SERV="protoc --rbp-rpc_out=${GOPATH}/src -I=./protoc-gen-rbp-rpc/proto -I=./proto -I=./gen_proto"

${PROTOC} rpc/user/user_profile_message.proto
${PROTO_SERV} rpc/user/user_profile_service.proto
${PROTOC} rpc/voice_lover/voice_lover_message.proto
${PROTO_SERV} rpc/voice_lover/voice_lover_main_service.proto
${PROTOC} rpc/voice_lover/voice_lover_common.proto
${PROTO_SERV} rpc/voice_lover/voice_lover_admin_service.proto
${PROTOC} rpc/voice_lover/voice_lover_admin_message.proto




