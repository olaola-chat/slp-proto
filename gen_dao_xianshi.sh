#!/bin/bash

cd `dirname $0`

rm -rf ./dao/xianshi
rm -rf ./gen_proto/db/xianshi
rm -rf ./gen_pb/db/xianshi

./gen_db.sh xianshi xs_user_profile
./gen_db.sh xianshi xs_chatroom
./gen_db.sh xianshi xs_chatroom_config

# ./gen_db.sh xianshi bbc_commodity_type_config
# ./gen_db.sh xianshi bbc_commodity_whitelist
# ./gen_db.sh xianshi bbc_murder_auto_question_category
# ./gen_db.sh xianshi bbc_room_black_list
# ./gen_db.sh xianshi bbc_sign_audit

./gen_db.sh xianshi xs_account_relationship
./gen_db.sh xianshi xs_wedding_album
