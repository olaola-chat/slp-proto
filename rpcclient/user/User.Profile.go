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

func (s *userProfile) NeedVerify(ctx context.Context, req *user.ReqNeedVerify) (*user.RepNeedVerify, error) {
	reply := &user.RepNeedVerify{}
	err := s.Call(ctx, "NeedVerify", req, reply)
	return reply, err
}

func (s *userProfile) Mget(ctx context.Context, req *user.ReqUserProfiles) (*user.RepUserProfiles, error) {
	reply := &user.RepUserProfiles{}
	err := s.Call(ctx, "Mget", req, reply)
	return reply, err
}

func (s *userProfile) Test(ctx context.Context, req *user.ReqUserProfiles) (*user.RepUserProfiles, error) {
	reply := &user.RepUserProfiles{}
	err := s.Call(ctx, "Test", req, reply)
	return reply, err
}

func (s *userProfile) SearchByName(ctx context.Context, req *user.ReqUserSearchName) (*user.RepUserSearchDefault, error) {
	reply := &user.RepUserSearchDefault{}
	err := s.Call(ctx, "SearchByName", req, reply)
	return reply, err
}
