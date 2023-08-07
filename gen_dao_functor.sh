#!/bin/bash

cd `dirname $0`

rm -rf ./dao/functor
rm -rf ./gen_proto/db/functor
rm -rf ./gen_pb/db/functor

./gen_db.sh functor voice_lover_album
./gen_db.sh functor voice_lover_album_comment
./gen_db.sh functor voice_lover_album_subject
./gen_db.sh functor voice_lover_audio
./gen_db.sh functor voice_lover_audio_album
./gen_db.sh functor voice_lover_audio_comment
./gen_db.sh functor voice_lover_subject
./gen_db.sh functor voice_lover_user_collect

# ./gen_db.sh xianshi bbc_commodity_type_config
# ./gen_db.sh xianshi bbc_commodity_whitelist
# ./gen_db.sh xianshi bbc_murder_auto_question_category
# ./gen_db.sh xianshi bbc_room_black_list
# ./gen_db.sh xianshi bbc_sign_audit
