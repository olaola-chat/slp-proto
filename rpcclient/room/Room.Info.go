package room

import (
	"context"

	"github.com/olaola-chat/slp-proto/rpcclient/base"

	"github.com/olaola-chat/slp-proto/gen_pb/db/xianshi"
	"github.com/olaola-chat/slp-proto/gen_pb/rpc/room"
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

func (s *roomInfo) InRoom(ctx context.Context, req *room.ReqUid) (*room.RepRid, error) {
	reply := &room.RepRid{}
	err := s.Call(ctx, "InRoom", req, reply)
	return reply, err
}

func (s *roomInfo) MgetInRoom(ctx context.Context, req *room.ReqUids) (*room.RepInRooms, error) {
	reply := &room.RepInRooms{}
	err := s.Call(ctx, "MgetInRoom", req, reply)
	return reply, err
}
