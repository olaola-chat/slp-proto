#! /bin/bash

cd `dirname $0`

if [ $# -lt 2 ];
then
        echo "USAGE: $0 dbname tableName"
        exit 1
fi

export GF_GCFG_FILE=config/config_cli.toml

genproto() {
    gf-cli gen pbentity -g $2 -t $1 -path gen_proto/db/$2

    if [ $? -ne 0 ]; then
        exit 1
    fi
}

gendao(){
    if [ "$2" = "xianshi" ]; then
        gf-cli gen dao -g default -t $1
    else
        gf-cli gen dao -g $2 -t $1
    fi

    if [ $? -ne 0 ]; then
        exit 1
    fi
}

sed "s/dbname/$1/g" config/daoindex.tpl > daoindex.tpl
sed "s/dbname/$1/g" config/modelindex.tpl > modelindex.tpl

genproto $2 $1
gendao $2 $1

rm daoindex.tpl modelindex.tpl

