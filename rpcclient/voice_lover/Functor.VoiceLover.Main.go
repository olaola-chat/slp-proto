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

func (s *voiceLoverMain) Post(ctx context.Context, req *voice_lover.ReqVoiceLoverPost) (*voice_lover.ResVoiceLoverBase, error) {
	reply := &voice_lover.ResVoiceLoverBase{}
	err := s.Call(ctx, "Post", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetRecAlbums(ctx context.Context, req *voice_lover.ReqGetRecAlbums) (*voice_lover.ResGetRecAlbums, error) {
	reply := &voice_lover.ResGetRecAlbums{}
	err := s.Call(ctx, "GetRecAlbums", req, reply)
	return reply, err
}
