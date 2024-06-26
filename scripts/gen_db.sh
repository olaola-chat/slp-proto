#! /bin/bash

cd `dirname $0`
cd ..

if [ $# -lt 2 ];
then
        echo "USAGE: $0 dbname tableName"
        exit 1
fi

export GF_GCFG_FILE=config/config_cli.toml

genproto() {
    gf-cli gen pbentity -g $2 -t $1 -path gen_proto/db/$2 -package "db.$2" -option "option go_package = \"github.com/olaola-chat/rbp-proto/gen_pb/db/$2\";"

    if [ $? -ne 0 ]; then
        echo "genproto failed"
        exit 1
    fi
}

gendao(){
    # if [ "$2" = "xianshi" ]; then
    #     gf-cli gen dao -g default -t $1
    # else
    #     gf-cli gen dao -g $2 -t $1
    # fi

    gf-cli gen dao -g $2 -t $1

    if [ $? -ne 0 ]; then
        echo "gendao failed"
        exit 1
    fi
}

gengo() {
    # protoc --go_out=. --proto_path=gen_proto/db/$2 entity_$1.proto
    protoc --go_out=. -I=./gen_proto db/$2/entity_$1.proto
    if [ $? -ne 0 ]; then
        echo "gengo failed"
        exit 1
    fi
}

inject() {
    protoc-go-inject-tag -input=gen_pb/db/$2/entity_$1.pb.go
    if [ $? -ne 0 ]; then
        echo "inject failed"
        exit 1
    fi
}

genproto $2 $1
gendao $2 $1
gengo $2 $1
inject $2 $1


