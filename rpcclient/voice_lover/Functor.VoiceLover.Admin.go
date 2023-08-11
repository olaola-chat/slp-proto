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

func (s *voiceLoverAdmin) CreateAlbum(ctx context.Context, req *voice_lover.ReqCreateAlbum) (*voice_lover.ResCreateAlbum, error) {
	reply := &voice_lover.ResCreateAlbum{}
	err := s.Call(ctx, "CreateAlbum", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) UpdateAlbum(ctx context.Context, req *voice_lover.ReqUpdateAlbum) (*voice_lover.ResUpdateAlbum, error) {
	reply := &voice_lover.ResUpdateAlbum{}
	err := s.Call(ctx, "UpdateAlbum", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) DelAlbum(ctx context.Context, req *voice_lover.ReqDelAlbum) (*voice_lover.ResDelAlbum, error) {
	reply := &voice_lover.ResDelAlbum{}
	err := s.Call(ctx, "DelAlbum", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) GetAlbumList(ctx context.Context, req *voice_lover.ReqGetAlbumList) (*voice_lover.ResGetAlbumList, error) {
	reply := &voice_lover.ResGetAlbumList{}
	err := s.Call(ctx, "GetAlbumList", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) GetAlbumDetail(ctx context.Context, req *voice_lover.ReqGetAlbumDetail) (*voice_lover.ResGetAlbumDetail, error) {
	reply := &voice_lover.ResGetAlbumDetail{}
	err := s.Call(ctx, "GetAlbumDetail", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AudioCollect(ctx context.Context, req *voice_lover.ReqAudioCollect) (*voice_lover.ResAudioCollect, error) {
	reply := &voice_lover.ResAudioCollect{}
	err := s.Call(ctx, "AudioCollect", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) CreateSubject(ctx context.Context, req *voice_lover.ReqCreateSubject) (*voice_lover.ResCreateSubject, error) {
	reply := &voice_lover.ResCreateSubject{}
	err := s.Call(ctx, "CreateSubject", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) DelSubject(ctx context.Context, req *voice_lover.ReqDelSubject) (*voice_lover.ResDelSubject, error) {
	reply := &voice_lover.ResDelSubject{}
	err := s.Call(ctx, "DelSubject", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) UpdateSubject(ctx context.Context, req *voice_lover.ReqUpdateSubject) (*voice_lover.ResUpdateSubject, error) {
	reply := &voice_lover.ResUpdateSubject{}
	err := s.Call(ctx, "UpdateSubject", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) GetSubjectDetail(ctx context.Context, req *voice_lover.ReqGetSubjectDetail) (*voice_lover.ResGetSubjectDetail, error) {
	reply := &voice_lover.ResGetSubjectDetail{}
	err := s.Call(ctx, "GetSubjectDetail", req, reply)
	return reply, err
}
