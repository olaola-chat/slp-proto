#!/bin/bash


cd `dirname $0`
cd ..

export PROTOC="protoc --go_out=${GOPATH}/src -I=./protoc-gen-slp-rpc/proto -I=./proto -I=./gen_proto"
export PROTO_SERV="protoc --slp-rpc_out=${GOPATH}/src -I=./protoc-gen-slp-rpc/proto -I=./proto -I=./gen_proto"

${PROTOC} rpc/user/auth_message.proto
${PROTO_SERV} rpc/user/auth_service.proto

${PROTOC} rpc/friends/friends_message.proto
${PROTO_SERV} rpc/friends/friends_service.proto

${PROTOC} rpc/user/user_profile_message.proto
${PROTO_SERV} rpc/user/user_profile_service.proto

${PROTOC} rpc/voice_lover/voice_lover_message.proto
${PROTO_SERV} rpc/voice_lover/voice_lover_main_service.proto
${PROTOC} rpc/voice_lover/voice_lover_common.proto
${PROTO_SERV} rpc/voice_lover/voice_lover_admin_service.proto
${PROTOC} rpc/voice_lover/voice_lover_admin_message.proto

${PROTOC} rpc/consume/consume_money_message.proto
${PROTO_SERV} rpc/consume/consume_money_service.proto

${PROTOC} rpc/room/room_info_message.proto
${PROTOC} rpc/room/room_package_message.proto
${PROTO_SERV} rpc/room/room_info_service.proto

${PROTOC} rpc/room/room_hot_message.proto
${PROTO_SERV} rpc/room/room_hot_service.proto

${PROTOC} rpc/activity/rank_message.proto
${PROTO_SERV} rpc/activity/rank_service.proto

${PROTOC} rpc/activity/send_gift_message.proto

${PROTOC} rpc/activity/qixi_message.proto
${PROTO_SERV} rpc/activity/qixi_service.proto
