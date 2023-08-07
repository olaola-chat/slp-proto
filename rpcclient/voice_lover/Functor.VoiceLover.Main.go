package voice_lover

import (
	"context"

	"github.com/olaola-chat/rbp-proto/rpcclient/base"

	"github.com/olaola-chat/rbp-proto/gen_pb/rpc/voice_lover"
)

// VoiceLoverMain *voiceLoverMain
var VoiceLoverMain = &voiceLoverMain{
	&base.Base{
		Name: "Functor.VoiceLover.Main",
	},
}

type voiceLoverMain struct {
	*base.Base
}

func (s *voiceLoverMain) Post(ctx context.Context, req *voice_lover.ReqVoiceLoverPost) (*voice_lover.ReqVoiceLoverPost, error) {
	reply := &voice_lover.ReqVoiceLoverPost{}
	err := s.Call(ctx, "Post", req, reply)
	return reply, err
}
