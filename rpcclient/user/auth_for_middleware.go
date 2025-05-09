package user

import (
	"context"

	"github.com/olaola-chat/slp-library/server/http/middleware"

	"github.com/olaola-chat/slp-proto/gen_pb/rpc/user"
)

func Auth(ctx context.Context, token string) (middleware.AuthUser, error) {
	req := &user.ReqUserAuth{Token: token}
	res, err := RbpUserAuth.Auth(ctx, req)
	if err != nil {
		return middleware.AuthUser{}, err
	}

	user := middleware.AuthUser{
		UID:      res.Uid,
		Time:     res.Time,
		AppID:    uint8(res.AppId),
		Salt:     res.Salt,
		Platform: res.Platform,
		Channel:  res.Channel,
	}

	return user, nil
}

func Auth2(ctx context.Context, token string) (middleware.AuthUser, error) {
	res, err := UserProfile.Auth(ctx, token)
	if err != nil {
		return middleware.AuthUser{}, err
	}

	user := middleware.AuthUser{
		UID:      res.Uid,
		Time:     res.Time,
		AppID:    uint8(res.AppId),
		Salt:     res.Salt,
		Platform: res.Platform,
		Channel:  res.Channel,
	}

	return user, nil
}
