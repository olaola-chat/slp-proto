package activity

import (
	"context"

	"github.com/olaola-chat/slp-proto/rpcclient/base"

	"github.com/olaola-chat/slp-proto/gen_pb/rpc/activitypb"
)

// RbpActivityRank *rbpActivityRank
var RbpActivityRank = &rbpActivityRank{
	&base.Base{
		Name: "Rbp.Activity.Rank",
	},
}

type rbpActivityRank struct {
	*base.Base
}

func (s *rbpActivityRank) Insert(ctx context.Context, req *activitypb.ReqRankInsert) (*activitypb.RepRankInsert, error) {
	reply := &activitypb.RepRankInsert{}
	err := s.Call(ctx, "Insert", req, reply)
	return reply, err
}

func (s *rbpActivityRank) IncrSimple(ctx context.Context, req *activitypb.ReqRankIncrSimple) (*activitypb.RepRankIncr, error) {
	reply := &activitypb.RepRankIncr{}
	err := s.Call(ctx, "IncrSimple", req, reply)
	return reply, err
}

func (s *rbpActivityRank) IncrComplex(ctx context.Context, req *activitypb.ReqRankIncrComplex) (*activitypb.RepRankIncr, error) {
	reply := &activitypb.RepRankIncr{}
	err := s.Call(ctx, "IncrComplex", req, reply)
	return reply, err
}

func (s *rbpActivityRank) GetList(ctx context.Context, req *activitypb.ReqRankGetList) (*activitypb.RepRankGetList, error) {
	reply := &activitypb.RepRankGetList{}
	err := s.Call(ctx, "GetList", req, reply)
	return reply, err
}
