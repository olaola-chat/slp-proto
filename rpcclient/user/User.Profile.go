package user

import (
	"context"

	"github.com/olaola-chat/rbp-proto/rpcclient/base"

	"github.com/olaola-chat/rbp-proto/gen_pb/rpc/user"
)

// UserProfile *userProfile
var UserProfile = &userProfile{
	&base.Base{
		Name: "User.Profile",
	},
}

type userProfile struct {
	*base.Base
}

func (s *userProfile) NeedVerify(ctx context.Context, req *user.ReqNeedVerify) (*user.ReqNeedVerify, error) {
	reply := &user.ReqNeedVerify{}
	err := s.Call(ctx, "NeedVerify", req, reply)
	return reply, err
}

func (s *userProfile) Mget(ctx context.Context, req *user.ReqUserProfiles) (*user.ReqUserProfiles, error) {
	reply := &user.ReqUserProfiles{}
	err := s.Call(ctx, "Mget", req, reply)
	return reply, err
}

func (s *userProfile) Test(ctx context.Context, req *user.ReqUserProfiles) (*user.ReqUserProfiles, error) {
	reply := &user.ReqUserProfiles{}
	err := s.Call(ctx, "Test", req, reply)
	return reply, err
}
