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

func (s *voiceLoverMain) Post(ctx context.Context, req *voice_lover.ReqPost) (*voice_lover.ResBase, error) {
	reply := &voice_lover.ResBase{}
	err := s.Call(ctx, "Post", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetRecAlbums(ctx context.Context, req *voice_lover.ReqGetRecAlbums) (*voice_lover.ResGetRecAlbums, error) {
	reply := &voice_lover.ResGetRecAlbums{}
	err := s.Call(ctx, "GetRecAlbums", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetAlbumList(ctx context.Context, req *voice_lover.ReqGetAlbumList) (*voice_lover.ResGetAlbumList, error) {
	reply := &voice_lover.ResGetAlbumList{}
	err := s.Call(ctx, "GetAlbumList", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetRecSubjects(ctx context.Context, req *voice_lover.ReqGetRecSubjects) (*voice_lover.ResGetRecSubjects, error) {
	reply := &voice_lover.ResGetRecSubjects{}
	err := s.Call(ctx, "GetRecSubjects", req, reply)
	return reply, err
}

func (s *voiceLoverMain) BatchGetAlbumAudioCount(ctx context.Context, req *voice_lover.ReqBatchGetAlbumAudioCount) (*voice_lover.ResBatchGetAlbumAudioCount, error) {
	reply := &voice_lover.ResBatchGetAlbumAudioCount{}
	err := s.Call(ctx, "BatchGetAlbumAudioCount", req, reply)
	return reply, err
}
