package user

import (
	"context"

	"github.com/olaola-chat/slp-proto/rpcclient/base"

	"github.com/olaola-chat/slp-proto/gen_pb/rpc/user"
)

// RbpUserAuth *rbpUserAuth
var RbpUserAuth = &rbpUserAuth{
	&base.Base{
		Name: "User.Profile",
	},
}

type rbpUserAuth struct {
	*base.Base
}

func (s *rbpUserAuth) Auth(ctx context.Context, req *user.ReqUserAuth) (*user.RepUserAuth, error) {
	reply := &user.RepUserAuth{}
	err := s.Call(ctx, "Auth", req, reply)
	return reply, err
}
