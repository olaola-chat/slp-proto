package room

import (
	"context"

	"github.com/olaola-chat/rbp-proto/rpcclient/base"

	"github.com/olaola-chat/rbp-proto/gen_pb/rpc/room"
)

// RoomInfo *roomInfo
var RoomInfo = &roomInfo{
	&base.Base{
		Name: "Room.Info",
	},
}

type roomInfo struct {
	*base.Base
}

func (s *roomInfo) InRoom(ctx context.Context, req *room.ReqUid) (*room.RepRid, error) {
	reply := &room.RepRid{}
	err := s.Call(ctx, "InRoom", req, reply)
	return reply, err
}
