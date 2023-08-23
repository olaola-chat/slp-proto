package room

import (
	"context"

	"github.com/olaola-chat/rbp-proto/rpcclient/base"

	"github.com/olaola-chat/rbp-proto/gen_pb/rpc/room"
)

// RoomHot *roomHot
var RoomHot = &roomHot{
	&base.Base{
		Name: "Room.Hot",
	},
}

type roomHot struct {
	*base.Base
}

func (s *roomHot) Get(ctx context.Context, req *room.ReqRoomHotGet) (*room.ReqRoomHotGet, error) {
	reply := &room.ReqRoomHotGet{}
	err := s.Call(ctx, "Get", req, reply)
	return reply, err
}

func (s *roomHot) GetBatch(ctx context.Context, req *room.ReqRoomHotGetBatch) (*room.ResRoomHotGetBatch, error) {
	reply := &room.ResRoomHotGetBatch{}
	err := s.Call(ctx, "GetBatch", req, reply)
	return reply, err
}
