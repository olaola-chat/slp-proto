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

func (s *voiceLoverMain) GetAlbumInfoById(ctx context.Context, req *voice_lover.ReqGetAlbumInfoById) (*voice_lover.ResGetAlbumInfoById, error) {
	reply := &voice_lover.ResGetAlbumInfoById{}
	err := s.Call(ctx, "GetAlbumInfoById", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetRecAlbums(ctx context.Context, req *voice_lover.ReqGetRecAlbums) (*voice_lover.ResGetRecAlbums, error) {
	reply := &voice_lover.ResGetRecAlbums{}
	err := s.Call(ctx, "GetRecAlbums", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetRecCommonAlbums(ctx context.Context, req *voice_lover.ReqGetRecCommonAlbums) (*voice_lover.ResGetRecAlbums, error) {
	reply := &voice_lover.ResGetRecAlbums{}
	err := s.Call(ctx, "GetRecCommonAlbums", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetAlbumsByPage(ctx context.Context, req *voice_lover.ReqGetAlbumsByPage) (*voice_lover.ResGetAlbumsByPage, error) {
	reply := &voice_lover.ResGetAlbumsByPage{}
	err := s.Call(ctx, "GetAlbumsByPage", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetSubjectAlbumsByPage(ctx context.Context, req *voice_lover.ReqGetSubjectAlbumsByPage) (*voice_lover.ResGetAlbumsByPage, error) {
	reply := &voice_lover.ResGetAlbumsByPage{}
	err := s.Call(ctx, "GetSubjectAlbumsByPage", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetRecSubjects(ctx context.Context, req *voice_lover.ReqGetRecSubjects) (*voice_lover.ResGetRecSubjects, error) {
	reply := &voice_lover.ResGetRecSubjects{}
	err := s.Call(ctx, "GetRecSubjects", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetAudioListByAlbumId(ctx context.Context, req *voice_lover.ReqGetAudioListByAlbumId) (*voice_lover.ResGetAudioListByAlbumId, error) {
	reply := &voice_lover.ResGetAudioListByAlbumId{}
	err := s.Call(ctx, "GetAudioListByAlbumId", req, reply)
	return reply, err
}

func (s *voiceLoverMain) BatchGetAlbumAudioCount(ctx context.Context, req *voice_lover.ReqBatchGetAlbumAudioCount) (*voice_lover.ResBatchGetAlbumAudioCount, error) {
	reply := &voice_lover.ResBatchGetAlbumAudioCount{}
	err := s.Call(ctx, "BatchGetAlbumAudioCount", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetAlbumCommentCount(ctx context.Context, req *voice_lover.ReqGetAlbumCommentCount) (*voice_lover.ResGetAlbumCommentCount, error) {
	reply := &voice_lover.ResGetAlbumCommentCount{}
	err := s.Call(ctx, "GetAlbumCommentCount", req, reply)
	return reply, err
}

func (s *voiceLoverMain) IsUserCollectAlbum(ctx context.Context, req *voice_lover.ReqIsUserCollectAlbum) (*voice_lover.ResIsUserCollectAlbum, error) {
	reply := &voice_lover.ResIsUserCollectAlbum{}
	err := s.Call(ctx, "IsUserCollectAlbum", req, reply)
	return reply, err
}

func (s *voiceLoverMain) IsUserCollectAlbums(ctx context.Context, req *voice_lover.ReqIsUserCollectAlbums) (*voice_lover.ResIsUserCollectAlbums, error) {
	reply := &voice_lover.ResIsUserCollectAlbums{}
	err := s.Call(ctx, "IsUserCollectAlbums", req, reply)
	return reply, err
}

func (s *voiceLoverMain) Collect(ctx context.Context, req *voice_lover.ReqCollect) (*voice_lover.ResCollect, error) {
	reply := &voice_lover.ResCollect{}
	err := s.Call(ctx, "Collect", req, reply)
	return reply, err
}

func (s *voiceLoverMain) SubmitAudioComment(ctx context.Context, req *voice_lover.ReqAudioSubmitComment) (*voice_lover.ResCommonPost, error) {
	reply := &voice_lover.ResCommonPost{}
	err := s.Call(ctx, "SubmitAudioComment", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetAudioCommentList(ctx context.Context, req *voice_lover.ReqGetAudioCommentList) (*voice_lover.ResCommentList, error) {
	reply := &voice_lover.ResCommentList{}
	err := s.Call(ctx, "GetAudioCommentList", req, reply)
	return reply, err
}

func (s *voiceLoverMain) SubmitAlbumComment(ctx context.Context, req *voice_lover.ReqAlbumSubmitComment) (*voice_lover.ResCommonPost, error) {
	reply := &voice_lover.ResCommonPost{}
	err := s.Call(ctx, "SubmitAlbumComment", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetAlbumCommentList(ctx context.Context, req *voice_lover.ReqGetAlbumCommentList) (*voice_lover.ResCommentList, error) {
	reply := &voice_lover.ResCommentList{}
	err := s.Call(ctx, "GetAlbumCommentList", req, reply)
	return reply, err
}

func (s *voiceLoverMain) GetAudioInfoById(ctx context.Context, req *voice_lover.ReqGetAudioDetail) (*voice_lover.ResGetAudioDetail, error) {
	reply := &voice_lover.ResGetAudioDetail{}
	err := s.Call(ctx, "GetAudioInfoById", req, reply)
	return reply, err
}
