package auth

import (
	"context"

	"github.com/olaola-chat/rbp-proto/rpcclient/base"

	"github.com/olaola-chat/rbp-proto/gen_pb/rpc/user/auth"
)

// RbpUserAuth *rbpUserAuth
var RbpUserAuth = &rbpUserAuth{
	&base.Base{
		Name: "Rbp.User.Auth",
	},
}

type rbpUserAuth struct {
	*base.Base
}

func (s *rbpUserAuth) Auth(ctx context.Context, req *auth.ReqUserAuth) (*auth.RepUserAuth, error) {
	reply := &auth.RepUserAuth{}
	err := s.Call(ctx, "Auth", req, reply)
	return reply, err
}
