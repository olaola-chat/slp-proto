package voice_lover

import (
	"context"

	"github.com/olaola-chat/rbp-proto/rpcclient/base"

	"github.com/olaola-chat/rbp-proto/gen_pb/rpc/voice_lover"
)

// VoiceLoverAdmin *voiceLoverAdmin
var VoiceLoverAdmin = &voiceLoverAdmin{
	&base.Base{
		Name: "Functor.VoiceLover.Admin",
	},
}

type voiceLoverAdmin struct {
	*base.Base
}

func (s *voiceLoverAdmin) GetAudioEdit(ctx context.Context, req *voice_lover.ReqGetAudioEdit) (*voice_lover.ResGetAudioEdit, error) {
	reply := &voice_lover.ResGetAudioEdit{}
	err := s.Call(ctx, "GetAudioEdit", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) GetAudioDetail(ctx context.Context, req *voice_lover.ReqGetAudioDetail) (*voice_lover.ResGetAudioDetail, error) {
	reply := &voice_lover.ResGetAudioDetail{}
	err := s.Call(ctx, "GetAudioDetail", req, reply)
	return reply, err
}
