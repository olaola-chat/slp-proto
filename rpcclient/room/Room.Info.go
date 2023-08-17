package room

import (
	"context"

	"github.com/olaola-chat/rbp-proto/rpcclient/base"

	"github.com/olaola-chat/rbp-proto/gen_pb/db/xianshi"
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

func (s *roomInfo) Get(ctx context.Context, req *room.ReqRoomInfo) (*xianshi.EntityXsChatroom, error) {
	reply := &xianshi.EntityXsChatroom{}
	err := s.Call(ctx, "Get", req, reply)
	return reply, err
}

func (s *roomInfo) Mget(ctx context.Context, req *room.ReqRoomInfos) (*room.RepRoomInfos, error) {
	reply := &room.RepRoomInfos{}
	err := s.Call(ctx, "Mget", req, reply)
	return reply, err
}
