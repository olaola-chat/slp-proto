package consume

import (
	"context"

	"github.com/olaola-chat/slp-proto/rpcclient/base"

	"github.com/olaola-chat/slp-proto/gen_pb/rpc/consume"
)

// ConsumeMoney *consumeMoney
var ConsumeMoney = &consumeMoney{
	&base.Base{
		Name: "Consume.Money",
	},
}

type consumeMoney struct {
	*base.Base
}

func (s *consumeMoney) GetRequestId(ctx context.Context, req *consume.GetRequestIdReq) (*consume.GetRequestIdRes, error) {
	reply := &consume.GetRequestIdRes{}
	err := s.Call(ctx, "GetRequestId", req, reply)
	return reply, err
}

func (s *consumeMoney) CommoditySend(ctx context.Context, req *consume.CommoditySendReq) (*consume.CommoditySendRes, error) {
	reply := &consume.CommoditySendRes{}
	err := s.Call(ctx, "CommoditySend", req, reply)
	return reply, err
}
