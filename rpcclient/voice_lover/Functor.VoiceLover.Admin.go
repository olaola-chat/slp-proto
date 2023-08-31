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

func (s *voiceLoverAdmin) GetSubjectList(ctx context.Context, req *voice_lover.ReqGetSubjectList) (*voice_lover.ResGetSubjectList, error) {
	reply := &voice_lover.ResGetSubjectList{}
	err := s.Call(ctx, "GetSubjectList", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) GetAlbumCollect(ctx context.Context, req *voice_lover.ReqGetAlbumCollect) (*voice_lover.ResGetAlbumCollect, error) {
	reply := &voice_lover.ResGetAlbumCollect{}
	err := s.Call(ctx, "GetAlbumCollect", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AlbumCollect(ctx context.Context, req *voice_lover.ReqAlbumCollect) (*voice_lover.ResAlbumCollect, error) {
	reply := &voice_lover.ResAlbumCollect{}
	err := s.Call(ctx, "AlbumCollect", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AlbumChoice(ctx context.Context, req *voice_lover.ReqAlbumChoice) (*voice_lover.ResAlbumChoice, error) {
	reply := &voice_lover.ResAlbumChoice{}
	err := s.Call(ctx, "AlbumChoice", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) GetAlbumChoice(ctx context.Context, req *voice_lover.ReqGetAlbumChoice) (*voice_lover.ResGetAlbumChoice, error) {
	reply := &voice_lover.ResGetAlbumChoice{}
	err := s.Call(ctx, "GetAlbumChoice", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) GetBannerList(ctx context.Context, req *voice_lover.ReqGetBannerList) (*voice_lover.ResGetBannerList, error) {
	reply := &voice_lover.ResGetBannerList{}
	err := s.Call(ctx, "GetBannerList", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) CreateBanner(ctx context.Context, req *voice_lover.ReqCreateBanner) (*voice_lover.ResCreateBanner, error) {
	reply := &voice_lover.ResCreateBanner{}
	err := s.Call(ctx, "CreateBanner", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) UpdateBanner(ctx context.Context, req *voice_lover.ReqUpdateBanner) (*voice_lover.ResUpdateBanner, error) {
	reply := &voice_lover.ResUpdateBanner{}
	err := s.Call(ctx, "UpdateBanner", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) GetBannerDetail(ctx context.Context, req *voice_lover.ReqGetBannerDetail) (*voice_lover.ResGetBannerDetail, error) {
	reply := &voice_lover.ResGetBannerDetail{}
	err := s.Call(ctx, "GetBannerDetail", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAudioList(ctx context.Context, req *voice_lover.ReqAdminAudioList) (*voice_lover.ResAdminAudioList, error) {
	reply := &voice_lover.ResAdminAudioList{}
	err := s.Call(ctx, "AdminAudioList", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAudioDetail(ctx context.Context, req *voice_lover.ReqAdminAudioDetail) (*voice_lover.ResAdminAudioDetail, error) {
	reply := &voice_lover.ResAdminAudioDetail{}
	err := s.Call(ctx, "AdminAudioDetail", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAudioUpdate(ctx context.Context, req *voice_lover.ReqAdminAudioUpdate) (*voice_lover.ResAdminAudioUpdate, error) {
	reply := &voice_lover.ResAdminAudioUpdate{}
	err := s.Call(ctx, "AdminAudioUpdate", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAudioAudit(ctx context.Context, req *voice_lover.ReqAdminAudioAudit) (*voice_lover.ResAdminAudioAudit, error) {
	reply := &voice_lover.ResAdminAudioAudit{}
	err := s.Call(ctx, "AdminAudioAudit", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAudioAuditReason(ctx context.Context, req *voice_lover.ReqAdminAudioAuditReason) (*voice_lover.ResAdminAudioAuditReason, error) {
	reply := &voice_lover.ResAdminAudioAuditReason{}
	err := s.Call(ctx, "AdminAudioAuditReason", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAlbumCreate(ctx context.Context, req *voice_lover.ReqAdminAlbumCreate) (*voice_lover.ResAdminAlbumCreate, error) {
	reply := &voice_lover.ResAdminAlbumCreate{}
	err := s.Call(ctx, "AdminAlbumCreate", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAlbumDel(ctx context.Context, req *voice_lover.ReqAdminAlbumDel) (*voice_lover.ResAdminAlbumDel, error) {
	reply := &voice_lover.ResAdminAlbumDel{}
	err := s.Call(ctx, "AdminAlbumDel", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAlbumUpdate(ctx context.Context, req *voice_lover.ReqAdminAlbumUpdate) (*voice_lover.ResAdminAlbumUpdate, error) {
	reply := &voice_lover.ResAdminAlbumUpdate{}
	err := s.Call(ctx, "AdminAlbumUpdate", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAlbumDetail(ctx context.Context, req *voice_lover.ReqAdminAlbumDetail) (*voice_lover.ResAdminAlbumDetail, error) {
	reply := &voice_lover.ResAdminAlbumDetail{}
	err := s.Call(ctx, "AdminAlbumDetail", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAlbumList(ctx context.Context, req *voice_lover.ReqAdminAlbumList) (*voice_lover.ResAdminAlbumList, error) {
	reply := &voice_lover.ResAdminAlbumList{}
	err := s.Call(ctx, "AdminAlbumList", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAudioCollectList(ctx context.Context, req *voice_lover.ReqAdminAudioCollectList) (*voice_lover.ResAdminAudioCollectList, error) {
	reply := &voice_lover.ResAdminAudioCollectList{}
	err := s.Call(ctx, "AdminAudioCollectList", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAudioCollect(ctx context.Context, req *voice_lover.ReqAdminAudioCollect) (*voice_lover.ResAdminAudioCollect, error) {
	reply := &voice_lover.ResAdminAudioCollect{}
	err := s.Call(ctx, "AdminAudioCollect", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminSubjectCreate(ctx context.Context, req *voice_lover.ReqAdminSubjectCreate) (*voice_lover.ResAdminSubjectCreate, error) {
	reply := &voice_lover.ResAdminSubjectCreate{}
	err := s.Call(ctx, "AdminSubjectCreate", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminSubjectUpdate(ctx context.Context, req *voice_lover.ReqAdminSubjectUpdate) (*voice_lover.ResAdminSubjectUpdate, error) {
	reply := &voice_lover.ResAdminSubjectUpdate{}
	err := s.Call(ctx, "AdminSubjectUpdate", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminSubjectDel(ctx context.Context, req *voice_lover.ReqAdminSubjectDel) (*voice_lover.ResAdminSubjectDel, error) {
	reply := &voice_lover.ResAdminSubjectDel{}
	err := s.Call(ctx, "AdminSubjectDel", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminSubjectList(ctx context.Context, req *voice_lover.ReqAdminSubjectList) (*voice_lover.ResAdminSubjectList, error) {
	reply := &voice_lover.ResAdminSubjectList{}
	err := s.Call(ctx, "AdminSubjectList", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAlbumCollect(ctx context.Context, req *voice_lover.ReqAdminAlbumCollect) (*voice_lover.ResAdminAlbumCollect, error) {
	reply := &voice_lover.ResAdminAlbumCollect{}
	err := s.Call(ctx, "AdminAlbumCollect", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAlbumChoice(ctx context.Context, req *voice_lover.ReqAdminAlbumChoice) (*voice_lover.ResAdminAlbumChoice, error) {
	reply := &voice_lover.ResAdminAlbumChoice{}
	err := s.Call(ctx, "AdminAlbumChoice", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAlbumCollectList(ctx context.Context, req *voice_lover.ReqAdminAlbumCollectList) (*voice_lover.ResAdminAlbumCollectList, error) {
	reply := &voice_lover.ResAdminAlbumCollectList{}
	err := s.Call(ctx, "AdminAlbumCollectList", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminSubjectDetail(ctx context.Context, req *voice_lover.ReqAdminSubjectDetail) (*voice_lover.ResAdminSubjectDetail, error) {
	reply := &voice_lover.ResAdminSubjectDetail{}
	err := s.Call(ctx, "AdminSubjectDetail", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminAlbumChoiceList(ctx context.Context, req *voice_lover.ReqAdminAlbumChoiceList) (*voice_lover.ResAdminAlbumChoiceList, error) {
	reply := &voice_lover.ResAdminAlbumChoiceList{}
	err := s.Call(ctx, "AdminAlbumChoiceList", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminBannerList(ctx context.Context, req *voice_lover.ReqAdminBannerList) (*voice_lover.ResAdminBannerList, error) {
	reply := &voice_lover.ResAdminBannerList{}
	err := s.Call(ctx, "AdminBannerList", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminBannerCreate(ctx context.Context, req *voice_lover.ReqAdminBannerCreate) (*voice_lover.ResAdminBannerCreate, error) {
	reply := &voice_lover.ResAdminBannerCreate{}
	err := s.Call(ctx, "AdminBannerCreate", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminBannerUpdate(ctx context.Context, req *voice_lover.ReqAdminBannerUpdate) (*voice_lover.ResAdminBannerUpdate, error) {
	reply := &voice_lover.ResAdminBannerUpdate{}
	err := s.Call(ctx, "AdminBannerUpdate", req, reply)
	return reply, err
}

func (s *voiceLoverAdmin) AdminBannerDetail(ctx context.Context, req *voice_lover.ReqAdminBannerDetail) (*voice_lover.ResAdminBannerDetail, error) {
	reply := &voice_lover.ResAdminBannerDetail{}
	err := s.Call(ctx, "AdminBannerDetail", req, reply)
	return reply, err
}
