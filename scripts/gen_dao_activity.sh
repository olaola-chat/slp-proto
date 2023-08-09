#!/bin/bash

cd `dirname $0`

./gen_db.sh activity act_experience_user_bind
./gen_db.sh activity act_experience_consume_statics

