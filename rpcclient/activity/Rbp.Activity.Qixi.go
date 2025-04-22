package activity

import (
	"context"

	"github.com/olaola-chat/slp-proto/rpcclient/base"

	"github.com/olaola-chat/slp-proto/gen_pb/rpc/activitypb"
)

// RbpActivityQixi *rbpActivityQixi
var RbpActivityQixi = &rbpActivityQixi{
	&base.Base{
		Name: "Rbp.Activity.Qixi",
	},
}

type rbpActivityQixi struct {
	*base.Base
}

func (s *rbpActivityQixi) RoomPackage(ctx context.Context, req *activitypb.ReqRoomPackage) (*activitypb.RepConsume, error) {
	reply := &activitypb.RepConsume{}
	err := s.Call(ctx, "RoomPackage", req, reply)
	return reply, err
}

func (s *rbpActivityQixi) ChatGift(ctx context.Context, req *activitypb.ReqChatGift) (*activitypb.RepConsume, error) {
	reply := &activitypb.RepConsume{}
	err := s.Call(ctx, "ChatGift", req, reply)
	return reply, err
}

func (s *rbpActivityQixi) GetMyData(ctx context.Context, req *activitypb.ReqQixiMyData) (*activitypb.RepQixiMyData, error) {
	reply := &activitypb.RepQixiMyData{}
	err := s.Call(ctx, "GetMyData", req, reply)
	return reply, err
}

func (s *rbpActivityQixi) GetAwardPool(ctx context.Context, req *activitypb.ReqQixiAwardPool) (*activitypb.RepQixiAwardPool, error) {
	reply := &activitypb.RepQixiAwardPool{}
	err := s.Call(ctx, "GetAwardPool", req, reply)
	return reply, err
}
