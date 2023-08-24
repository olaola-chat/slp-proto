package friends

import (
	"context"

	"github.com/olaola-chat/rbp-proto/rpcclient/base"

	"github.com/olaola-chat/rbp-proto/gen_pb/rpc/friends"
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

func (s *friend) Count(ctx context.Context, req *friends.ReqFriendCount) (*friends.RepFriendCount, error) {
	reply := &friends.RepFriendCount{}
	err := s.Call(ctx, "Count", req, reply)
	return reply, err
}
