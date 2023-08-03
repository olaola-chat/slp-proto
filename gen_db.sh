#! /bin/bash

cd `dirname $0`

export GF_GCFG_FILE=config/config_cli.toml

genproto() {
    if [ -z "$1" ] 
    then
        exit 1
    fi

    db=$2

    if [ "$db" ]
    then
        gf-cli gen pbentity -g $db -t $1
    else
        gf-cli gen pbentity -g default -t $1
    fi
    if [ $? -ne 0 ]; then
        exit 1
    fi
}

gendao(){
    if [ -z "$1" ] 
    then
        exit 1
    fi

    name=$1
    db=$2

    if [ "$db" ] 
    then
        gf-cli gen dao -g $db -t $name
    else
        gf-cli gen dao -g default -t $name
    fi
    if [ $? -ne 0 ]; then
        exit 1
    fi
}

sed "s/dbname/$1/g" config/daoindex.tpl > daoindex.tpl
sed "s/dbname/$1/g" config/modelindex.tpl > modelindex.tpl

genproto $1 $2
gendao $1 $2

rm daoindex.tpl modelindex.tpl

