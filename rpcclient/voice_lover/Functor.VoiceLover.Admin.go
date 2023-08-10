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

func (s *voiceLoverAdmin) GetAudioDetail(ctx context.Context, req *voice_lover.ReqGetAudioDetail) (*voice_lover.ResGetAudioDetail, error) {
	reply := &voice_lover.ResGetAudioDetail{}
	err := s.Call(ctx, "GetAudioDetail", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) UpdateAudio(ctx context.Context, req *voice_lover.ReqUpdateAudio) (*voice_lover.ResBase, error) {
	reply := &voice_lover.ResBase{}
	err := s.Call(ctx, "UpdateAudio", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AuditAudio(ctx context.Context, req *voice_lover.ReqAuditAudio) (*voice_lover.ResAuditAudio, error) {
	reply := &voice_lover.ResAuditAudio{}
	err := s.Call(ctx, "AuditAudio", req, reply)
	return reply, err
}
