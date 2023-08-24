package friend

import (
	"context"

	"github.com/olaola-chat/rbp-proto/rpcclient/base"

	"github.com/olaola-chat/rbp-proto/gen_pb/rpc/friend"
)

// Friend *friend
var Friend = &friend{
	&base.Base{
		Name: "Friend",
	},
}

type friend struct {
	*base.Base
}

func (s *friend) Count(ctx context.Context, req *friend.ReqFriendCount) (*friend.RepFriendCount, error) {
	reply := &friend.RepFriendCount{}
	err := s.Call(ctx, "Count", req, reply)
	return reply, err
}
