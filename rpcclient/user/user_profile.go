package user

import (
	"context"

	"github.com/olaola-chat/rbp-proto/gen_pb/db/xianshi"
	"github.com/olaola-chat/rbp-proto/gen_pb/rpc/user"
)

func MgetMap(ctx context.Context, uids []uint32, fields []string) (map[uint32]*xianshi.EntityXsUserProfile, error) {
	if len(uids) == 0 {
		return nil, nil
	}

	req := &user.ReqUserProfiles{
		Uids: uids,
	}

	hasUid := false
	for _, field := range fields {
		if field == "uid" {
			hasUid = true
		}
		req.Fields = append(req.Fields, field)
	}

	//必须带uid
	if len(fields) > 0 && !hasUid {
		req.Fields = append(req.Fields, "uid")
	}

	data, err := UserProfile.Mget(ctx, req)
	if err != nil || data == nil {
		return nil, err
	}

	ret := make(map[uint32]*xianshi.EntityXsUserProfile, len(data.Data))

	for _, v := range data.Data {
		ret[v.Uid] = v
	}

	return ret, nil
}
