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

func (s *voiceLoverAdmin) ReqCreateAlbum(ctx context.Context, req *voice_lover.ReqCreateAlbum) (*voice_lover.ResCreateAlbum, error) {
	reply := &voice_lover.ResCreateAlbum{}
	err := s.Call(ctx, "ReqCreateAlbum", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) ReqUpdateAlbum(ctx context.Context, req *voice_lover.ReqUpdateAlbum) (*voice_lover.ResUpdateAlbum, error) {
	reply := &voice_lover.ResUpdateAlbum{}
	err := s.Call(ctx, "ReqUpdateAlbum", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) ReqDelAlbum(ctx context.Context, req *voice_lover.ReqDelAlbum) (*voice_lover.ResDelAlbum, error) {
	reply := &voice_lover.ResDelAlbum{}
	err := s.Call(ctx, "ReqDelAlbum", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) ReqGetAlbumList(ctx context.Context, req *voice_lover.ReqGetAlbumList) (*voice_lover.ResGetAlbumList, error) {
	reply := &voice_lover.ResGetAlbumList{}
	err := s.Call(ctx, "ReqGetAlbumList", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) ReqGetAlbumDetail(ctx context.Context, req *voice_lover.ReqGetAlbumDetail) (*voice_lover.ResGetAlbumDetail, error) {
	reply := &voice_lover.ResGetAlbumDetail{}
	err := s.Call(ctx, "ReqGetAlbumDetail", req, reply)
	return reply, err
}
