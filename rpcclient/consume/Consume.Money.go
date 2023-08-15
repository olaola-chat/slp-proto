package consume

import (
	"context"

	"github.com/olaola-chat/rbp-proto/rpcclient/base"

	"github.com/olaola-chat/rbp-proto/gen_pb/rpc/consume"
)

// UserProfile *userProfile
var UserProfile = &userProfile{
	&base.Base{
		Name: "Consume.Money",
	},
}

type userProfile struct {
	*base.Base
}

func (s *userProfile) GetRequestId(ctx context.Context, req *consume.GetRequestIdReq) (*consume.GetRequestIdRes, error) {
	reply := &consume.GetRequestIdRes{}
	err := s.Call(ctx, "GetRequestId", req, reply)
	return reply, err
}

func (s *userProfile) CommoditySend(ctx context.Context, req *consume.CommoditySendReq) (*consume.CommoditySendRes, error) {
	reply := &consume.CommoditySendRes{}
	err := s.Call(ctx, "CommoditySend", req, reply)
	return reply, err
}
